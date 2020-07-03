import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import CoreInput from "sharedComponents/core/Input/CoreInput";
import Dropdown, { DropdownOption } from "sharedComponents/core/Input/Dropdown";
import Button from "sharedComponents/core/Input/Button";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex',
        flexDirection: 'column',
        padding: '65px',
        borderRadius: '14px',
        boxShadow: '0 1px 7px 3px rgba(0,0,0,0.11)'
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'space-between'
    },
    header: {
        fontSize: theme.fontSizes.extraLargeHeading,
        fontWeight: '900',
        marginBottom: '15px'
    },
    subtitle: {
        fontSize: theme.fontSizes.extraLarge,
        color: theme.colors.textGrey,
        marginBottom: '30px',
        maxWidth: '288px'
    },
    blue: {
        cursor: 'pointer',
        color: theme.colors.navyBlue
    },
    border: {
        width: '60px',
        borderBottom: ['3px', 'solid', theme.colors.navyBlue],
        marginBottom: '45px',
    },
    input: {
        fontSize: theme.fontSizes.default,
        padding: '11px',
        border: ['2px', 'solid', theme.colors.borderGrey],
        borderRadius: '6px',
        marginBottom: '30px'
    },
    marginLeft: {
        marginLeft: '27px'
    },
    dropdown: {
        marginBottom: '30px'
    },
    dropdownBox: {
        border: ['2px', 'solid', theme.colors.borderGrey],
        borderRadius: '6px'
    },
    dropdownText: {
        fontSize: theme.fontSizes.default,
        fontWeight: '400',
        color: theme.colors.textGrey
    },
    last: {
        resize: 'none',
        justifyContent: 'flex-start'
    },
    button: {
        height: '52px',
        width: '163px',
        fontSize: theme.fontSizes.extraLarge,
        fontWeight: '800',
        boxShadow: '0 1px 4px 0 rgba(0,0,0,0.43)'
    }
}));

export type ContactDetails = {
    firstName: string;
    lastName: string;
    companyName: string;
    email: string;
    type: string;
    extra: string;
}

type Props = {
    contactDetails: ContactDetails;
    onHelp: () => void;
    onSubmit: () => void;
    className?: string;
};

function ContactPage({ contactDetails, onHelp, onSubmit, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    // Form Data
    const [firstName, setFirstName] = React.useState("");
    const [lastName, setLastName] = React.useState("");
    const [companyName, setCompanyName] = React.useState("");
    const [email, setEmail] = React.useState("");
    const [type, setType] = React.useState({id: 0, title: 'How can we help you today?'});
    const [extra, setExtra] = React.useState("");

    React.useEffect(() => {
        contactDetails = {
            firstName,
            lastName,
            companyName,
            email,
            type: type.title !== 'How can we help you today?' ? type.title : "",
            extra,
        };
    });

  return (
    <form className={classNames(classes.root, className)}>   
        <div className={classes.header}>Get in touch</div>
        <div className={classes.subtitle}>Explore our <span className={classes.blue} onClick={onHelp}>Help Guides</span> or contact our team below</div>
        <div className={classes.border} />
        <div className={classes.row}>
            <CoreInput
                placeholder="First Name"
                type="text"
                onChange={(text: string) => setFirstName(text)}
                value={firstName}
                className={classes.input}
            />

            <CoreInput
                placeholder="Last Name"
                type="text"
                onChange={(text: string) => setLastName(text)}
                value={lastName}
                className={classNames(classes.input, classes.marginLeft)}
            />
        </div>

        <CoreInput
            placeholder="Company Name"
            type="text"
            onChange={(text: string) => setCompanyName(text)}
            value={companyName}
            className={classes.input}
        />

        <CoreInput
            placeholder="Email Address"
            type="text"
            onChange={(text: string) => setEmail(text)}
            value={email}
            className={classes.input}
        />

        <Dropdown
            placeholder="How can we help you today?"
            options={[{id: 0, title: 'Assistance Type'}]}
            selected={type}
            setSelected={(selected: DropdownOption) => setType(selected)}
            fontStyle={classes.dropdownText}
            boxStyle={classes.dropdownBox}
            className={classes.dropdown}
        />

        <textarea
            placeholder="Go into as much detail as you like, we're all ears!"
            rows={8}
            onChange={(event: any) => setExtra(event.target.value)}
            value={extra}
            className={classNames(classes.input, classes.last)}
        />
        <Button
            archetype="submit"
            className={classes.button}
            onClick={onSubmit}
        >
            Submit
        </Button>
    </form>
  );
}

export default ContactPage;