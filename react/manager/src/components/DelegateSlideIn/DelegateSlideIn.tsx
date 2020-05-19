import * as React from "react";
import SideModal from "components/core/SideModal";
import Tabs, { TabContent, Body, Heading, Footer } from "components/core/SideModal/Tabs";
import FancyInput from "components/LoginDialogue/FancyInput";
import Button from "sharedComponents/core/Button";
import Icon from "sharedComponents/core/Icon";

type Props = {
  isOpen: boolean;
  onClose: () => void;
  submitDelegate: (
    profileUrl: string,
    firstName: string,
    lastName: string,
    jobTitle: string,
    email: string,
    phone: string,
    ttcId: string,
    errorCallback: (err : string) => void
    ) => void
};

const delegateDetails: TabContent[] = [
  {
    key: "Delegate Details",
    component: ({ state, setState, setTab, closeModal}) => (
      <>
        <Body>
          <Heading>
            Personal
          </Heading>
          <div style= {{ display: "flex", marginTop: 25, marginBottom: 35}}>
            <div style={{ display: "flex", 
                        flexDirection: "column",
                        width: "30%",
                        alignItems: "center",
                        justifyContent: "center", 
                        marginRight: "5%"}}>
              <Icon
                name="EditPicPencil"
                size={30}
                style={{marginBottom: -30, marginRight: -65, zIndex: 15}}
              />
              <Icon
                name="EditDelegateProfilePic_Default"
                pointer
                size={100}
                onClick={() => alert("ratta tatta!")}
              />
              <Button archetype="default" style={{marginTop: 10}}>
                Edit Photo
              </Button>
            </div>
            <div style={{ display: "flex", flexDirection: "column"}}>
              <div style={{ display: "flex", flexDirection: "row", marginBottom: 10}}>
                <div style={{ marginRight: 20 }}>
                  <FancyInput
                    label="First Name"
                    placeholder="e.g. John"
                    onChange={(firstName) => 
                      setState((s: object) => ({ ...s, firstName}))
                    }
                  />
                </div>
                <FancyInput
                  label="Last Name"
                  placeholder="e.g. Smith"
                  onChange={(lastName) => 
                    setState((s: object) => ({ ...s, lastName}))
                  }
                />
              </div>
              <FancyInput
                label="Job Title"
                placeholder="e.g. Mechanical Engineer"
                onChange={(jobTitle) => 
                  setState((s: object) => ({ ...s, jobTitle}))
                }
              />
            </div>
          </div>
          <Heading>
            Contact
          </Heading>
          <div style={{ display: "flex" , marginTop: 5, marginBottom: 35}}>
            <div style={{ marginRight: "7.5%", width: "40%"}}>
              <FancyInput
                label="Email" 
                type="email"
                placeholder="e.g. example@fedex.com"
                onChange={(email) => 
                  setState((s: object) => ({ ...s, email}))
                }
              />
            </div>
            <div style={{ width: "40%"}}>
              <FancyInput
                label="Phone number"
                type="tel"
                placeholder="e.g. +44 1234 567890"
                onChange={(phone) => 
                  setState((s: object) => ({ ...s, phone}))
                }
              />
            </div>
          </div>
          <Heading>
            TTC ID
          </Heading>

          <div style={{ width: "50%"}}>
            <FancyInput
              label=""
              onChange={(ttcId) => 
                setState((s: object) => ({ ...s, ttcId}))
              }
            />
          </div>
        </Body>
        <Footer>
          <div />
          <div style={{ display: "flex" }}>
            <Button archetype="default" onClick={() => closeModal()}>
              Cancel
            </Button>
            <Button
              archetype="submit"
              onClick={() => state.submitDelegate(
                state.profileUrl,
                state.firstName,
                state.lastName,
                state.jobTitle,
                state.email,
                state.phone,
                state.ttcId,
                (err : string) => alert(err)
              )}
              style={{ marginLeft: 20 }}
            >
              Confirm {"&"} Save
            </Button>
          </div>
        </Footer>
      </>
    )
  }
];

const DelegateSlideIn = ({ 
   isOpen,
   onClose,
   submitDelegate
  }: Props) => {

  return (
    <SideModal title="Add New Delegate" closeModal={onClose} isOpen={isOpen}>
      <Tabs
        content={delegateDetails}
        closeModal={onClose}
        initialState={{
          profileUrl: "https://i.imgur.com/C0RGBYP.jpg",
          firstName: "",
          lastName: "",
          jobTitle: "",
          email: "",
          phone: "",
          ttcId: "",
          onSubmit: submitDelegate
        }}
        />
    </SideModal>
  );
}

export default DelegateSlideIn;