import * as React from 'react';
import SideModal from './SideModal';
import { withKnobs, boolean } from '@storybook/addon-knobs';
import Button from 'sharedComponents/core/Button';
import Tabs, {
  TabContent,
  Footer,
  Body,
  Heading,
  LargeText,
  TermsBox,
  Text,
  CourseList,
  CurrentTotal,
  Payment
} from './Tabs';
import SearchableDropdown, {
  CourseCategory,
  Course
} from '../SearchableDropdown';
import { ResultItem } from 'components/UserSearch';
import MultiUserSearch from 'components/UserSearch/MultiUserSearch';
import Checkbox from '../Checkbox';

export default {
  title: 'Core/Course Management',
  decorators: [withKnobs]
};

const trainingInitialValue = [
  {
    label:
      'I Fred Eccs can confirm that the above delegates have undergone the necessary 5-year background checks for this Training Programme',
    checked: false
  }
];

const ToBInitialValue = [
  {
    label:
      'I Fred Eccs can confirm that I have read and understood the TTC Terms of Business',
    checked: false
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
      700
    )
  );
};

const items: ResultItem[] = [
  {
    key: 'Jim Smith',
    value: 'uuid-1'
  },
  {
    key: 'Bruce Willis',
    value: 'uuid-2'
  },
  {
    key: 'Tony Stark',
    value: 'uuid-3'
  }
];

const userSearchFunction = async (query: string) => {
  return await new Promise<ResultItem[]>((resolve) =>
    setTimeout(() => resolve(items), 700)
  );
};

const tabs: TabContent[] = [
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
            selected={state.course ? [state.course] : []}
            searchQuery={courseSearchFunction}
            setSelected={(selected) =>
              setState((s: object) => ({ ...s, courses: selected }))
            }
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
      <div style={{ margin: '30px 40px' }}>
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
      </div>
    )
  }
];

const tabs2: TabContent[] = [
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
      <div style={{ margin: '30px 40px' }}>
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
      </div>
    )
  }
];

export const multipleUsers = () => {
  return (
    <SideModal
      isOpen={boolean('Open', true)}
      title="Course Management"
      closeModal={() => alert('Close Function called')}
    >
      <Tabs
        content={tabs}
        closeModal={() => alert('Close Function called')}
        initialState={{
          course: undefined,
          users: [undefined],
          training: trainingInitialValue,
          ToB: ToBInitialValue
        }}
      />
    </SideModal>
  );
};

export const singleUser = () => {
  return (
    <SideModal
      isOpen={boolean('Open', true)}
      title="Course Management"
      closeModal={() => alert('Close Function called')}
    >
      <Tabs
        content={tabs2}
        closeModal={() => alert('Close Function called')}
        initialState={{
          courses: [],
          training: trainingInitialValue,
          ToB: ToBInitialValue
        }}
      />
    </SideModal>
  );
};
