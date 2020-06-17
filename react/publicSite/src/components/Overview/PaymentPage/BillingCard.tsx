import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import CoreInput from "sharedComponents/core/Input/CoreInput";
import Dropdown from "sharedComponents/core/Input/Dropdown";
import { DropdownOption } from "sharedComponents/core/Input/Dropdown/Dropdown";
import CheckboxSingle from "sharedComponents/core/Input/CheckboxSingle";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex',
        flexDirection: 'column',
        padding: '50px 60px',
        borderRadius: '14px',
        boxShadow: '0 1px 7px 3px rgba(0,0,0,0.11)'
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'space-between'
    },
    header: {
        fontSize: '40px',
        fontWeight: '900',
        marginBottom: '50px'
    },
    input: {
        fontSize: theme.fontSizes.xSmallHeading,
        padding: '15px',
        border: ['2px', 'solid', theme.colors.borderGrey],
        borderRadius: '6px',
        marginBottom: '21px'
    },
    marginLeft: {
        marginLeft: '27px'
    },
    dropdown: {
        marginBottom: '41px',
        flex: 1.05
    },
    dropdownBox: {
        height: '56px',
        border: ['2px', 'solid', theme.colors.borderGrey],
        borderRadius: '6px'
    },
    dropdownText: {
        fontSize: theme.fontSizes.xSmallHeading,
        color: theme.colors.textGrey
    },
    checkbox: {
        flex: 1.05
    },
    checkboxText: {
        fontSize: theme.fontSizes.extraLarge,
        color: theme.colors.textGrey,
        marginLeft: '23px'
    }
}));

export type BillingDetails = {
    email: string;
    companyName: string;
    phoneNumber: string;
    firstName: string;
    lastName: string;
    adressOne: string;
    adressTwo: string;
    city: string;
    postcode: string;
    country: string;
    contact: boolean;
}

type Props = {
    billingDetails: BillingDetails;
    className?: string;
};

function BillingCard({ billingDetails, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    // Form Data
    const [email, setEmail] = React.useState("");
    const [companyName, setCompanyName] = React.useState("");
    const [phoneNumber, setPhoneNumber] = React.useState("");
    const [firstName, setFirstName] = React.useState("");
    const [lastName, setLastName] = React.useState("");
    const [adressOne, setAddressOne] = React.useState("");
    const [adressTwo, setAddressTwo] = React.useState("");
    const [city, setCity] = React.useState("");
    const [postcode, setPostcode] = React.useState("");
    const [country, setCountry] = React.useState({id: 0, title: 'Country'});
    const [contact, setContact] = React.useState(false);

    React.useEffect(() => {
        billingDetails = {
            email,
            companyName,
            phoneNumber,
            firstName,
            lastName,
            adressOne,
            adressTwo,
            city,
            postcode,
            country: country.title !== 'Country' ? country.title : "",
            contact,
        };
    });

  return (
    <form className={classNames(classes.root, className)}>   
        <div className={classes.header}>Billing Address</div>

        <CoreInput
            placeholder="Company Email Address"
            type="text"
            onChange={(text: string) => setEmail(text)}
            value={email}
            className={classes.input}
        />

        <CoreInput
            placeholder="Company Name"
            type="text"
            onChange={(text: string) => setCompanyName(text)}
            value={companyName}
            className={classes.input}
        />

        <CoreInput
            placeholder="Contact Telephone Number"
            type="text"
            onChange={(text: string) => setPhoneNumber(text)}
            value={phoneNumber}
            className={classes.input}
        />
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
            placeholder="Address Line 1"
            type="text"
            onChange={(text: string) => setAddressOne(text)}
            value={adressOne}
            className={classes.input}
        />

        <CoreInput
            placeholder="Address Line 1"
            type="text"
            onChange={(text: string) => setAddressTwo(text)}
            value={adressTwo}
            className={classes.input}
        />

        <div className={classes.row}>
            <CoreInput
                placeholder="City"
                type="text"
                onChange={(text: string) => setCity(text)}
                value={city}
                className={classes.input}
            />

            <Dropdown
                placeholder="Country"
                options={[{id: 0, title: 'U.K'}]}
                selected={country}
                setSelected={(selected: DropdownOption) => setCountry(selected)}
                fontStyle={classes.dropdownText}
                boxStyle={classes.dropdownBox}
                className={classNames(classes.dropdown, classes.marginLeft)}
            />
        </div>

        <div className={classes.row}>
            <CoreInput
                placeholder="Postcode"
                type="text"
                onChange={(text: string) => setPostcode(text)}
                value={postcode}
                className={classes.input}
            />

            <CheckboxSingle
                className={classNames(classes.checkbox, classes.marginLeft)}
                fontStyle={classes.checkboxText}
                size={30}
                onChange={() => setContact(!contact)}
                label="By checking this box you confirm you are happy for our team to contact you during the registration period"
            />
        </div> 
    </form>
  );
}

export default BillingCard;