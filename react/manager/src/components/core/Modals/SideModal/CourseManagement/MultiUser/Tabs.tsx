import * as React from 'react';
import { fetchQuery, graphql } from 'relay-runtime';
import Tabs, {
  TabContent,
  Body,
  Heading,
  LargeText,
  Footer,
  CurrentTotal,
  TermsBox,
  Text,
  CourseList,
  Payment
} from 'components/core/Modals/SideModal/Tabs';
import SearchableDropdown, {
  CourseCategory,
  Course
} from 'components/core/Input/SearchableDropdown';
import MultiUserSearch from 'components/UserSearch/MultiUserSearch';
import Button from 'sharedComponents/core/Input/Button';
import Checkbox from 'components/core/Input/Checkbox';
import { ResultItem } from 'components/UserSearch';
import environment from 'api/environment';
import { TabsDelegateQueryResponse } from './__generated__/TabsDelegateQuery.graphql';
import { TabsCoursesQueryResponse } from './__generated__/TabsCoursesQuery.graphql';

const items: ResultItem[] = [
  {
    uuid: '',
    key: 'Jim Smith',
    value: 'fedex-jimsmith'
  },
  {
    uuid: '',
    key: 'Bruce Willis',
    value: 'fedex-brucewillis'
  },
  {
    uuid: '',
    key: 'Tony Stark',
    value: 'fedex-tonystark'
  }
];

const categories: CourseCategory[] = [
  {
    title: 'Dangerous Goods - Air',
    courses: [
      {
        id: 1,
        name: 'Cargo Operative Screener (COS) – VC, HS, XRY, ETD',
        price: 200,
        trainingReq: true
      }
    ]
  },
  {
    title: 'Known Consignor',
    courses: [
      {
        id: 3,
        name: 'Known Consignor Responsible Person',
        price: 70,
        trainingReq: true
      },
      {
        id: 4,
        name: 'Known Consignor (Modules 1-7)',
        price: 55,
        trainingReq: false
      },
      {
        id: 5,
        name: 'Manual Handling Awareness',
        price: 25,
        trainingReq: false
      },
      {
        id: 6,
        name: 'Fire Safety Awareness',
        price: 25,
        trainingReq: false
      },
      {
        id: 7,
        name: 'Noise and Vibration Awareness',
        price: 25,
        trainingReq: false
      },
      {
        id: 8,
        name: 'Seat Belt Misuse Awareness',
        price: 300,
        trainingReq: true
      }
    ]
  }
];

const userSearchFunction = async (text: string) => {
  const query = graphql`
    query TabsDelegateQuery($name: String!) {
      delegates(filter: { name: $name }, page: { limit: 8 }) {
        edges {
          uuid
          TTC_ID
          firstName
          lastName
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
  )) as TabsDelegateQueryResponse;

  if (!data || !data.delegates || !data.delegates.edges) {
    console.error('Could not get data', data);
    return [];
  }

  const results = data.delegates?.edges.map((delegate) => ({
    uuid: delegate?.uuid ?? '',
    key: `${delegate?.firstName} ${delegate?.lastName}`,
    value: delegate?.TTC_ID ?? ''
  }));

  return results;
};

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
          <Heading>
            Start a Quick Booking{' '}
            {state.users.filter(Boolean).length > 0 && (
              <span style={{ color: '#1081AA' }}>
                for {state.users.filter(Boolean).length} Delegates
              </span>
            )}
          </Heading>
          <LargeText>Select your Course</LargeText>
          <SearchableDropdown
            selected={state.course}
            searchQuery={courseSearchFunction}
            setSelected={(selected) => {
              console.log(state.course);
              console.log(selected);
              setState((s: object) => ({ ...s, course: selected }));
            }}
          />
          <MultiUserSearch
            searchFunction={userSearchFunction}
            users={state.users}
            setUsers={(users) => setState((s: object) => ({ ...s, users }))}
            style={{ marginTop: 25 }}
          />
        </Body>
        <Footer>
          <CurrentTotal
            total={
              state.courses &&
              state.courses[0].price * state.users.filter(Boolean).length
            }
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
    component: ({ state, setState, closeModal, setTab }) => (
      <>
        <Body>
          <Heading>Terms of Business</Heading>
          <LargeText>
            In order to book {state.users.filter(Boolean).length} Delegates onto
            this Course, please refer to and confirm you have read our Terms of
            Business below.
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
              pariatur saepe minus? Voluptas sint voluptatum nihil, suscipit sed
              eaque rem porro at officiis eos voluptatibus, ullam cupiditate?
              Nobis porro adipisci animi, vitae ex vel?
            </Text>
          </TermsBox>
          {state.course && state.course.trainingReq > 0 && (
            <>
              <Heading>Background Check</Heading>
              <LargeText>
                The following Courses require you to validate that the above
                delegates have recently completed a 5-year background check to
                comply with Civil Aviation Authority standards.
              </LargeText>
              <CourseList courses={[state.course]} />
              <Checkbox
                boxes={state.training}
                setBoxes={(training) =>
                  setState((s: object) => ({ ...s, training }))
                }
              />
            </>
          )}
          <Checkbox
            boxes={state.ToB}
            setBoxes={(ToB) => setState((s: object) => ({ ...s, ToB }))}
          />
        </Body>
        <Footer>
          <CurrentTotal
            total={
              state.courses &&
              state.courses[0].price * state.users.filter(Boolean).length
            }
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
    )
  },
  {
    key: 'Payment',
    component: ({ state }) => (
      <Body>
        <Heading>Payment</Heading>
        <LargeText>
          You have opted to book {state.users.filter(Boolean).length} Delegates
          on the following Course
        </LargeText>
        <Payment
          courses={[
            {
              id: 1,
              name:
                'Cargo Operative Screener (COS) Recurrent – VC, HS, XRY, ETD',
              price: 55,
              sku: '082739428374',
              qty: 2,
              subtotal: 110
            },
            {
              id: 2,
              name:
                'Cargo Operative Screener (COS) Recurrent – VC, HS, XRY, ETD',
              price: 55,
              sku: '082739428374',
              qty: 3,
              subtotal: 165
            }
          ]}
        />
      </Body>
    )
  }
];
