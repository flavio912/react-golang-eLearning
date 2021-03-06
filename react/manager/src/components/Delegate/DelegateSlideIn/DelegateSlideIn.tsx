import * as React from 'react';
import SideModal from 'sharedComponents/SideModal/SideModal';
import Tabs, {
  TabContent,
  Body,
  Heading,
  Footer
} from 'sharedComponents/SideModal/Tabs';
import Button from 'sharedComponents/core/Input/Button';
import Icon from 'sharedComponents/core/Icon';
import EasyInput from 'components/core/Input/EasyInput';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import { CreateDelegate, UpdateDelegate } from './mutations';
import Checkbox from 'components/core/Input/Checkbox';
import CheckboxSingle from 'components/core/Input/CheckboxSingle';
import { check } from 'prettier';
import { mutations_CreateDelegateMutationResponse } from './__generated__/mutations_CreateDelegateMutation.graphql';
import IdentTag from 'sharedComponents/IdentTag';
import LabelledCard from 'sharedComponents/core/Cards/LabelledCard';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import { mutations_UpdateDelegateMutationResponse } from './__generated__/mutations_UpdateDelegateMutation.graphql';

const useStyles = createUseStyles((theme: Theme) => ({
  personalContainer: {
    display: 'flex',
    marginTop: theme.spacing(2),
    marginBottom: 35
  },
  heading: {
    marginBottom: theme.spacing(2),
    color: theme.colors.primaryBlack,
    fontSize: theme.fontSizes.heading,
    fontWeight: 800
  },
  profilePicContainer: {
    display: 'flex',
    flexDirection: 'column',
    width: '30%',
    alignItems: 'center',
    justifyContent: 'center',
    marginRight: theme.spacing(3) //"5%"
  },
  editPicPencilIcon: {
    marginBottom: -30,
    marginRight: -65,
    zIndex: 15
  },
  editPhotoBtn: {
    marginTop: theme.spacing(1)
  },
  personalDetailContainer: {
    display: 'flex',
    flexDirection: 'column'
  },
  namesContainer: {
    display: 'flex',
    flexDirection: 'row',
    marginBottom: theme.spacing(2)
  },
  firstName: {
    marginRight: theme.spacing(2)
  },
  contactContainer: {
    display: 'flex',
    marginTop: theme.spacing(0),
    marginBottom: 35
  },
  emailInput: {
    marginRight: '7.5%',
    width: '40%'
  },
  phoneInput: {
    width: '40%'
  },
  ttcInput: {
    marginBottom: theme.spacing(2)
  },
  ttcLabel: {
    fontSize: theme.fontSizes.heading,
    fontWeight: 800,
    marginLeft: 1,
    marginBottom: theme.spacing(1),
    flex: 1
  },
  ttcId: {
    flex: 2,
    border: `1px solid ${theme.colors.borderBlack}`,
    borderRadius: 5,
    padding: '20px 10px'
  },
  footerBtnsContainer: {
    display: 'flex'
  },
  submitBtn: {
    marginLeft: theme.spacing(2)
  },
  infoText: {
    fontSize: theme.fontSizes.default,
    color: theme.colors.textGrey,
    marginBottom: theme.spacing(1)
  },
  passwordText: {
    color: theme.colors.primaryBlack
  }
}));

export type DelegateInfo = {
  uuid: string;
  firstName: string;
  lastName: string;
  jobTitle: string;
  email: string;
  phone: string;
  ttcId: string;
  generatedPassword: string | null | undefined;
  generatePassword: boolean;
};

type Props = {
  isOpen: boolean;
  delegate?: DelegateInfo;
  onClose: () => void;
};

const delegateDetails: TabContent[] = [
  {
    key: 'Delegate Details',
    component: ({ state, setState, setTab, closeModal }) => {
      const theme = useTheme();
      const classes = useStyles({ theme });

      return (
        <>
          <Body>
            <span className={classes.heading}>Personal</span>
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
                  onClick={() => alert('ratta tatta!')}
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
                      value={state.firstName}
                      onChange={(firstName) =>
                        setState((s: object) => ({ ...s, firstName }))
                      }
                    />
                  </div>
                  <EasyInput
                    label="Last Name"
                    placeholder="e.g. Smith"
                    value={state.lastName}
                    onChange={(lastName) =>
                      setState((s: object) => ({ ...s, lastName }))
                    }
                  />
                </div>
                <EasyInput
                  label="Job Title"
                  placeholder="e.g. Mechanical Engineer"
                  value={state.jobTitle}
                  onChange={(jobTitle) =>
                    setState((s: object) => ({ ...s, jobTitle }))
                  }
                />
              </div>
            </div>
            <span className={classes.heading}>Contact</span>
            <div className={classes.contactContainer}>
              <div className={classes.emailInput}>
                <EasyInput
                  label="Email"
                  type="email"
                  placeholder="e.g. example@fedex.com"
                  value={state.email}
                  onChange={(email) =>
                    setState((s: object) => ({ ...s, email }))
                  }
                />
              </div>
              <div className={classes.phoneInput}>
                <EasyInput
                  label="Phone number"
                  type="tel"
                  value={state.phone}
                  placeholder="e.g. +44 1234 567890"
                  onChange={(phone) =>
                    setState((s: object) => ({ ...s, phone }))
                  }
                />
              </div>
            </div>
            {/* <Heading>
            TTC ID
          </Heading> */}
            {state.update ? (
              <div className={classes.ttcInput}>
                <div className={classes.ttcLabel}>TTC ID</div>
                <IdentTag ident={state.ttcId} />
              </div>
            ) : (
              <div>
                <span className={classes.heading}>Generate password</span>
                <p className={classes.infoText}>
                  If unchecked, a link will be sent to the email above to set
                  the user's password
                </p>
                <CheckboxSingle
                  label={'Generate a password for the new delegate'}
                  onChange={(checked: boolean) => {
                    setState((s: object) => ({
                      ...s,
                      generatePassword: checked
                    }));
                  }}
                />
              </div>
            )}
            {state.generatedPassword && (
              <LabelledCard label={'Generated password'}>
                <p className={classes.infoText}>
                  You will not be able to see this again, please take a note of
                  it:
                </p>
                <p className={classes.passwordText}>
                  {state.generatedPassword}
                </p>
              </LabelledCard>
            )}
          </Body>
          <Footer>
            <div />
            <div className={classes.footerBtnsContainer}>
              <Button archetype="default" onClick={() => closeModal()}>
                Cancel
              </Button>
              <Button
                archetype="submit"
                onClick={() => {
                  const delegate = {
                    uuid: state.uuid,
                    firstName: state.firstName,
                    lastName: state.lastName,
                    jobTitle: state.jobTitle,
                    email: state.email,
                    phone: state.phone,
                    generatePassword: state.generatePassword
                  };
                  state.onSubmit(delegate,
                    (resp : DelegateInfo) => setState(
                      (s : object) => ({
                        ...s,
                        firstName: resp.firstName,
                        lastName: resp.lastName,
                        jobTitle: resp.jobTitle,
                        email: resp.email,
                        phone: resp.phone,
                        generatedPassword: resp.generatedPassword,
                        ttcId: resp.ttcId,
                        update: true
                      })
                    )
                  );
                }
                }
                className={classes.submitBtn}
              >
                {state.update ? 'Update' : 'Confirm & Save'}
              </Button>
            </div>
          </Footer>
        </>
      );
    }
  }
];

const DelegateSlideIn = ({ isOpen, delegate, onClose }: Props) => {
  const inputDel = delegate;
  if (delegate === undefined) {
    delegate = {
      uuid: '',
      firstName: '',
      lastName: '',
      jobTitle: '',
      email: '',
      phone: '',
      ttcId: '',
      generatePassword: false,
      generatedPassword: '',
    };
  }

  const update = inputDel !== undefined;

  const createDelegate = (delegate: DelegateInfo, stateUpdate: (resp: DelegateInfo) => void) => 
    CreateDelegate(
      {
        firstName: delegate.firstName,
        lastName: delegate.lastName,
        jobTitle: delegate.jobTitle,
        email: delegate.email,
        phone: delegate.phone,
      },
      delegate.generatePassword ?? false,
      () => {},
      (resp: mutations_CreateDelegateMutationResponse) => stateUpdate(
        {
          uuid: delegate.uuid,
          firstName: resp.createDelegate?.delegate.firstName ?? '',
          lastName: resp.createDelegate?.delegate.lastName ?? '',
          jobTitle: resp.createDelegate?.delegate.jobTitle ?? '',
          email: resp.createDelegate?.delegate.email ?? '',
          phone: resp.createDelegate?.delegate.telephone ?? '',
          generatePassword: delegate.generatePassword,
          ttcId: resp.createDelegate?.delegate.TTC_ID ?? '',
          generatedPassword: resp.createDelegate?.generatedPassword
        }
      )
    );

  const updateDelegate = (delegate: DelegateInfo, stateUpdate: (resp: DelegateInfo) => void) => 
    UpdateDelegate(
      {
        uuid: delegate.uuid,
        firstName: delegate.firstName,
        lastName: delegate.lastName,
        jobTitle: delegate.jobTitle,
        email: delegate.email,
        phone: delegate.phone,
      },
      () => {},
      (resp: mutations_UpdateDelegateMutationResponse) => stateUpdate(
        {
          uuid: delegate.uuid,
          firstName: resp.updateDelegate?.firstName ?? '',
          lastName: resp.updateDelegate?.lastName ?? '',
          jobTitle: resp.updateDelegate?.jobTitle ?? '',
          email: resp.updateDelegate?.email ?? '',
          phone: resp.updateDelegate?.telephone ?? '',
          generatePassword:
            delegate.generatePassword,
          ttcId: resp.updateDelegate?.TTC_ID ?? '',
          generatedPassword: delegate.generatedPassword
        }
      )
    );

  const onSubmit = update
    ? updateDelegate
    : createDelegate;

  return (
    <SideModal
      title={
        update
          ? `${delegate.firstName} ${delegate.lastName}`
          : 'Add New Delegate'
      }
      closeModal={onClose}
      isOpen={isOpen}
    >
      <Tabs
        content={delegateDetails}
        closeModal={onClose}
        initialState={{
          profileUrl: 'https://i.imgur.com/C0RGBYP.jpg',
          uuid: delegate.uuid,
          firstName: delegate.firstName,
          lastName: delegate.lastName,
          jobTitle: delegate.jobTitle,
          email: delegate.email,
          phone: delegate.phone,
          ttcId: delegate.ttcId,
          update: update,
          onSubmit: onSubmit
        }}
      />
    </SideModal>
  );
};

export default DelegateSlideIn;
