import * as React from 'react';
import Tabs, {
  TabContent,
  Body,
  Heading,
  LargeText,
  Footer,
  CurrentTotal,
  TermsBox,
  Text,
  Payment
} from 'sharedComponents/SideModal/Tabs';
import SearchableDropdown, {
  CourseCategory,
  Course
} from 'sharedComponents/core/Input/SearchableDropdown';
import Button from 'sharedComponents/core/Input/Button';
import Checkbox from 'sharedComponents/core/Input/Checkbox';
import { graphql, fetchQuery } from 'react-relay';
import environment from 'api/environment';
import { TabsCoursesQueryResponse } from './__generated__/TabsCoursesQuery.graphql';

const courseSearchFunction = async (text: string) => {
  const query = graphql`
    query TabsCoursesQuery($name: String!) {
      courses(filter: { name: $name }, page: { limit: 60 }) {
        edges {
          ident: id
          name
          price
          category {
            uuid
            name
            color
          }
        }
      }
    }
  `;

  const variables = {
    name: text
  };

  const data = (await fetchQuery(
    environment,
    query,
    variables
  )) as TabsCoursesQueryResponse;

  if (!data || !data.courses || !data.courses.edges) {
    console.error('Could not get course data', data);
    return [];
  }

  const results: CourseCategory[] = [];

  // Get categories
  const categories = {};
  const uuidToTitle = {};
  data.courses?.edges.forEach((course) => {
    const courseItem = {
      id: course?.ident,
      name: course?.name,
      price: course?.price,
      trainingReq: false
    };

    const index = course?.category?.uuid;
    if (!index) return;

    const name = course?.category?.name;
    if (name == undefined) return;

    uuidToTitle[index] = name;

    if (categories.hasOwnProperty(index)) {
      categories[index].push(courseItem);
    } else {
      categories[index] = [courseItem];
    }
  });

  for (const key in categories) {
    results.push({
      title: uuidToTitle[key],
      courses: categories[key]
    });
  }

  return results;
};

export const tabList: TabContent[] = [
  {
    key: 'Courses',
    component: ({ state, setState, setTab, closeModal }) => (
      <>
        <Body>
          <Heading>Book a new course</Heading>
          <LargeText>You have selected the following courses</LargeText>
          <SearchableDropdown
            multiselect
            selected={state.courses}
            searchQuery={courseSearchFunction}
            setSelected={(courses) =>
              setState((s: object) => ({ ...s, courses }))
            }
          />
        </Body>
        <Footer>
          <CurrentTotal
            total={state.courses
              .map(({ price }: { price: number }) => price)
              .reduce((a: number, b: number) => a + b, 0)}
          />
          <div style={{ display: 'flex' }}>
            <Button archetype="default" onClick={() => closeModal()}>
              Cancel
            </Button>
            <Button
              archetype="submit"
              onClick={() => setTab('Terms of Business')}
              style={{ marginLeft: 20 }}
            >
              Continue to Terms
            </Button>
          </div>
        </Footer>
      </>
    )
  },
  {
    key: 'Terms of Business',
    component: ({ state, setState, closeModal, setTab }) => {
      return (
        <>
          <Body>
            <Heading>Terms of Business</Heading>
            <LargeText>
              In order to book John onto these Courses, please refer to and
              confirm you have read our Terms of Business below.
            </LargeText>
            <TermsBox title="TTC Hub - Terms of Business">
              <Text>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Aut
                deleniti nam porro optio! Nam quo enim ipsum eligendi in nihil
                perferendis, eaque voluptatem esse dolore quaerat laboriosam rem
                ipsa reprehenderit.
              </Text>
              <Text>
                Corporis voluptate molestias saepe placeat consequatur, pariatur
                recusandae ducimus at suscipit corrupti cupiditate, harum sint
                libero laudantium quaerat ipsum? Sint, ut nisi.
              </Text>
              <Text>
                Ipsam perferendis, id nobis autem, veniam porro magnam cum ex
                expedita in placeat nemo asperiores aliquam sequi illo aliquid
                pariatur saepe minus? Voluptas sint voluptatum nihil, suscipit
                sed eaque rem porro at officiis eos voluptatibus, ullam
                cupiditate? Nobis porro adipisci animi, vitae ex vel?
              </Text>
            </TermsBox>
            <Checkbox
              boxes={state.ToB}
              setBoxes={(ToB) => setState((s: object) => ({ ...s, ToB }))}
            />
          </Body>
          <Footer>
            <CurrentTotal
              total={state.courses
                .map(({ price }: { price: number }) => price)
                .reduce((a: number, b: number) => a + b, 0)}
            />
            <div style={{ display: 'flex' }}>
              <Button archetype="default" onClick={() => closeModal()}>
                Cancel
              </Button>
              <Button
                archetype="submit"
                onClick={() => setTab('Payment')}
                style={{ marginLeft: 20 }}
              >
                Continue to Payment
              </Button>
            </div>
          </Footer>
        </>
      );
    }
  },
  {
    key: 'Payment',
    component: ({ state }) => (
      <Body>
        <Payment
          courses={state.courses}
          userUUIDs={[]}
          isContract={false}
          onPurchase={() => false}
          onSuccess={() => {}}
          onError={() => {}}
        />
      </Body>
    )
  }
];
