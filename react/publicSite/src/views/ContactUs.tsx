import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import ContactPage, { ContactDetails } from "components/Overview/ContactPage";

const useStyles = createUseStyles((theme: Theme) => ({
    contactRoot: {
        width: '100%'
      },
      centerer: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center'
      },
      centered: {
        display: 'flex',
        justifyContent: 'space-around',
        width: '100%',
        maxWidth: '622px',
        margin: '40px 0 120px 0'
      },
      page: {
        marginTop: '100px',
        width: '100%',
        maxWidth: '622px;'
      },
      infoBox: {
        maxWidth: '175px;'
      },
      heading: {
        fontSize: theme.fontSizes.large,
        color: theme.colors.primaryBlack,
        fontWeight: '800',
        marginBottom: '15px'
      },
      text: {
        fontSize: theme.fontSizes.small,
        color: theme.colors.textGrey,
        marginBottom: '15px'
      },
      link: {
        cursor: 'pointer',
        color: theme.colors.navyBlue2
      }
}));

const emptyContactDetails: ContactDetails = {
    firstName: "",
    lastName: "",
    companyName: "",
    email: "",
    type: "",
    extra: "",
}

type Props = {};

function ContactUs(props: Props) {
  const classes = useStyles();
  return (
      <div className={classes.contactRoot}>
          <div className={classes.centerer}>
                <ContactPage
                    className={classes.page}
                    contactDetails={emptyContactDetails}
                    onHelp={() => console.log('Help')}
                    onSubmit={() => console.log('Submit')}
                />
                <div className={classes.centered}>
                    <div className={classes.infoBox}>
                        <div className={classes.heading}>TTC Hub</div>
                        <div className={classes.text}>168 Wey House, 15 Church St Weybridge, KT13 8NA  </div>
                    </div>
                    <div className={classes.infoBox}>
                        <div className={classes.heading}>Prefer Email?</div>
                        <div className={classes.text}>Contact us direct on <span className={classes.link}>admin@ttc-hub.com</span></div>
                    </div>
                </div>
          </div>
      </div>
  );
}

export default ContactUs;