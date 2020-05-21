import * as React from "react";
import SideModal from "components/core/SideModal";
import Tabs, { TabContent, Body, Heading, Footer } from "components/core/SideModal/Tabs";
import Button from "sharedComponents/core/Button";
import Icon from "sharedComponents/core/Icon";
import EasyInput from "components/core/EasyInput";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  personalContainer:{
    display: "flex",
    marginTop: theme.spacing(2),
    marginBottom: 35
  },
  heading:{
    marginBottom: theme.spacing(2),
    color: theme.colors.primaryBlack,
    fontSize: theme.fontSizes.heading,
    fontWeight: 800,
  },
  profilePicContainer:{
    display: "flex",
    flexDirection: "column",
    width: "30%",
    alignItems: "center",
    justifyContent: "center", 
    marginRight: theme.spacing(3) //"5%"
  },
  editPicPencilIcon:{
    marginBottom: -30,
    marginRight: -65,
    zIndex: 15
  },
  editPhotoBtn:{
    marginTop: theme.spacing(1),
  },
  personalDetailContainer:{
    display: "flex",
    flexDirection: "column"
  },
  namesContainer:{
    display: "flex",
    flexDirection: "row",
    marginBottom: theme.spacing(2)
  },
  firstName:{
    marginRight: theme.spacing(2)
  },
  contactContainer:{
    display: "flex" ,
    marginTop: theme.spacing(0),
    marginBottom: 35
  },
  emailInput:{
    marginRight: "7.5%",
    width: "40%"
  },
  phoneInput:{
    width: "40%"
  },
  ttcInput:{
    width: "50%"
  },
  ttcLabel:{
    fontSize: theme.fontSizes.heading,
    fontWeight: 800,
    marginLeft: 1
  },
  footerBtnsContainer:{
    display: "flex",
  },
  submitBtn:{
    marginLeft: theme.spacing(2)
  }
}));


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
    component: ({ state, setState, setTab, closeModal}) => {
      const theme = useTheme();
      const classes = useStyles({ theme });

      return (
      <>
        <Body>
          <span className={classes.heading}>
            Personal
          </span>
          <div className={classes.personalContainer}>
            <div className={classes.profilePicContainer}>
              <Icon
                name="EditPicPencil"
                size={30}
                className={classes.editPicPencilIcon}
              />
              <Icon
                name="EditDelegateProfilePic_Default"
                pointer
                size={100}
                onClick={() => alert("ratta tatta!")}
              />
              <Button archetype="grey" className={classes.editPhotoBtn}>
                Edit Photo
              </Button>
            </div>
            <div className={classes.personalDetailContainer}>
              <div className={classes.namesContainer}>
                <div className={classes.firstName}>
                  <EasyInput
                    label="First Name"
                    placeholder="e.g. John"
                    onChange={(firstName) => 
                      setState((s: object) => ({ ...s, firstName}))
                    }
                  />
                </div>
                <EasyInput
                  label="Last Name"
                  placeholder="e.g. Smith"
                  onChange={(lastName) => 
                    setState((s: object) => ({ ...s, lastName}))
                  }
                />
              </div>
              <EasyInput
                label="Job Title"
                placeholder="e.g. Mechanical Engineer"
                onChange={(jobTitle) => 
                  setState((s: object) => ({ ...s, jobTitle}))
                }
              />
            </div>
          </div>
          <span className={classes.heading}>
            Contact
          </span>
          <div className={classes.contactContainer}>
            <div className={classes.emailInput}>
              <EasyInput
                label="Email" 
                type="email"
                placeholder="e.g. example@fedex.com"
                onChange={(email) => 
                  setState((s: object) => ({ ...s, email}))
                }
              />
            </div>
            <div className={classes.phoneInput}>
              <EasyInput
                label="Phone number"
                type="tel"
                placeholder="e.g. +44 1234 567890"
                onChange={(phone) => 
                  setState((s: object) => ({ ...s, phone}))
                }
              />
            </div>
          </div>
          {/* <Heading>
            TTC ID
          </Heading> */}
          <div className={classes.ttcInput}>
            <EasyInput
              label="TTC ID"
              labelClassName={classes.ttcLabel}
              onChange={(ttcId) => 
                setState((s: object) => ({ ...s, ttcId}))
              }
            />
          </div>
        </Body>
        <Footer>
          <div />
          <div className={classes.footerBtnsContainer}>
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
              className={classes.submitBtn}
            >
              Confirm {"&"} Save
            </Button>
          </div>
        </Footer>
      </>
      );
    }
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