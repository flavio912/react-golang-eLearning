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
  CourseList,
  Payment
} from 'sharedComponents/SideModal/Tabs';
import SearchableDropdown, {
  CourseCategory,
  Course
} from 'components/core/Input/SearchableDropdown';
import MultiUserSearch from 'components/UserSearch/MultiUserSearch';
import Button from 'sharedComponents/core/Input/Button';
import Checkbox from 'components/core/Input/Checkbox';
import { ResultItem } from 'components/UserSearch';

const items: ResultItem[] = [
  {
    uuid: '',
    key: 'Jim Smith',
    value: 'uuid-1'
  },
  {
    uuid: '',
    key: 'Bruce Willis',
    value: 'uuid-2'
  },
  {
    uuid: '',
    key: 'Tony Stark',
    value: 'uuid-3'
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
      },
      {
        id: 2,
        name: 'Cargo Operative Screener (COS) Recurrent – VC, HS, XRY, ETD',
        price: 65,
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

const userSearchFunction = async (query: string) => {
  return await new Promise<ResultItem[]>((resolve) =>
    setTimeout(() => resolve(items), 400)
  );
};

const courseSearchFunction = async (query: string) => {
  return await new Promise<CourseCategory[]>((resolve) =>
    setTimeout(
      () =>
        resolve(
          categories
            .map(({ title, courses }) => ({
              title,
              courses: courses.filter(({ name }) =>
                name.toLocaleLowerCase().includes(query.toLocaleLowerCase())
              )
            }))
            .filter(({ courses }) => courses.length > 0)
        ),
      400
    )
  );
};

export const tabList: TabContent[] = [
  {
    key: 'Courses',
    component: ({ state, setState, setTab, closeModal }) => (
      <>
        <Body>
          <Heading>Book John's first Course</Heading>
          <LargeText>Book John on Course(s)</LargeText>
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
      const trainingReqCourses = state.courses.filter(
        (course: Course) => !!course.trainingReq
      );
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
            {trainingReqCourses.length > 0 && (
              <>
                <Heading>Background Check</Heading>
                <LargeText>
                  The following Course require you to validate that the above
                  delegates have recently completed a 5-year background check to
                  comply with Civil Aviation Authority standards.
                </LargeText>
                <CourseList courses={trainingReqCourses} />
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
