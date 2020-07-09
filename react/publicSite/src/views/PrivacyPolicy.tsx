import * as React from "react";
import classNames from 'classnames';
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import TOSHeader from "components/Overview/TOSPage/TOSHeader";
import TOSContents, { LinkDetails } from "components/Overview/TOSPage/TOSContents";
import PageMargin from "components/core/PageMargin";

const useStyles = createUseStyles((theme: Theme) => ({
    privacyRoot: {
      width: '100%',
      backgroundColor: theme.colors.primaryWhite
    },
    contents: {
      display: 'flex',
      flex: 1,
      marginRight: '20px'
    },
    heading: {
      marginBottom: '65px'
    },
    pageHeading: {
      fontSize: theme.fontSizes.extraLargeHeading,
      color: theme.colors.primaryBlack,
      fontWeight: '800',
    },
    pageSubheading: {
      fontSize: theme.fontSizes.xSmallHeading,
      color: theme.colors.textGrey,
      fontWeight: '200',
      margin: '15px 0'
    },
    lastUpdated: {
      fontSize: theme.fontSizes.default,
      color: theme.colors.textGrey,
      fontWeight: '200',
      paddingBottom: '10px',
      borderBottom: ['1px', 'solid', '#D8D8D8'],
      marginBottom: '25px'
    },
    textHeading: {
      fontSize: theme.fontSizes.heading,
      color: theme.colors.primaryBlack,
      fontWeight: '800',
      margin: '30px 0 10px 0'
    },
    text: {
      fontSize: theme.fontSizes.large,
      color: '#404349',
      fontWeight: '200',
      margin: '17px 0'
    },
    row: {
      display: 'flex',
      alignItems: 'center',
    },
    bold: {
      fontWeight: 'bold'
    },
    marginBottom: {
      marginBottom: '100px'
    },
    bulletFilled: {
      minHeight: '8px',
      minWidth: '8px',
      borderRadius: '8px',
      backgroundColor: theme.colors.textBlue,
      margin: '20px'
    },
    bulletEmpty: {
      minHeight: '8px',
      minWidth: '8px',
      borderRadius: '8px',
      border: ['1px', 'solid', theme.colors.textBlue],
      margin: '20px 20px 20px 40px'
  }
}));

const defaultLinks: LinkDetails[] = [
    { title: 'Terms of Service', link: '/' },
    { title: 'Privacy Policy', link: '/' },
    { title: 'Security Policy', link: '/' },
    { title: 'GDPR', link: '/' },
    { title: 'Cookie Policy', link: '/' },
    { title: 'List of Sub-processors', link: '/' }
]

type Props = {};

function PrivacyPolicy(props: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [selected, setSelected] = React.useState(defaultLinks[0]);

  return (
      <div className={classes.privacyRoot}>
          <TOSHeader
            className={classes.heading}
            title="Policies and Procedures"
            subtitle="We’re committed to keeping your data secure, your private information private, and being transparent about our practices as a business."
          />
          <PageMargin>
            <TOSContents
              className={classes.contents}
              links={defaultLinks}
              selected={selected}
              setSelected={setSelected}
            />
            <div style={{ flex: 3 }}>
              <div className={classes.pageHeading}>Terms of Service</div>
              <div className={classes.pageSubheading}>Policies that apply to your account and use of this site</div>
              <div className={classes.lastUpdated}>Last revised on February 10 2020, and effective as of that date</div>
              <div className={classes.text}>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in<br/><br/>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</div>
              <div className={classes.textHeading}>Your Account</div>
              <div className={classes.row}>
                  <div className={classes.bulletFilled}/>
                  <div className={classes.text}>
                      <span className={classes.bold}>Eligibility. </span>
                      Your account must be registered by a human. Accounts registered by "bots" or other automated methods are not permitted. Additionally, you must be 18 years of age or older.
                  </div>
              </div>
              <div className={classes.row}>
                  <div className={classes.bulletFilled}/>
                  <div className={classes.text}>
                      <span className={classes.bold}>Registration Information. </span>
                      You must provide a valid permanent email address, along with any other information required by Help Scout during the registration process. One person or legal entity may not maintain more than one free account.
                  </div>
              </div>
              <div className={classes.row}>
                  <div className={classes.bulletFilled}/>
                  <div className={classes.text}>
                      <span className={classes.bold}>Password. </span>
                      You are responsible for maintaining the security of your account and password. We will not be liable for any loss or damage from your failure to comply with this security obligation. Personally identifiable information submitted by you will be subject to our Privacy Policy.
                  </div>
              </div>
              <div className={classes.row}>
                  <div className={classes.bulletFilled}/>
                  <div className={classes.text}>
                      <span className={classes.bold}>Restrictions. </span>
                      You may not use the TTC Hub Service for any illegal or unauthorised purpose. You must not, in the use of the TTC Hub Service, violate any laws in your jurisdiction, including, among other things, by:
                  </div>
              </div>
              <div className={classes.row}>
                  <div className={classes.bulletEmpty}/>
                  <div className={classes.text}>
                      distributing any virus, time bomb, trap door, or other harmful or disruptive computer code, mechanism or program;
                  </div>
              </div>
              <div className={classes.row}>
                  <div className={classes.bulletEmpty}/>
                  <div className={classes.text}>
                      covering or obscuring any notice, legend, warning or banner contained on the Help Scout Service;
                  </div>
              </div>
              <div className={classes.row}>
                  <div className={classes.bulletEmpty}/>
                  <div className={classes.text}>
                      interfering with or circumventing any security feature of the Help Scout Service or any feature that restricts or enforces 	limitations on use of or access to the Help Scout Service;
                  </div>
              </div>
              <div className={classes.row}>
                  <div className={classes.bulletEmpty}/>
                  <div className={classes.text}>
                      infringing or violating the rights of any other party, including without limitation any intellectual property rights, including 	copyright 	laws, or rights of privacy or publicity
                  </div>
              </div>
              <div className={classes.row}>
                  <div className={classes.bulletEmpty}/>
                  <div className={classes.text}>
                      being obscene, offensive, pornographic, fraudulent, deceptive, defamatory, threatening, harassing, abusive, slanderous, 	hateful, 	or causing embarrassment to any other person as determined by Help Scout in its sole discretion; or
                  </div>
              </div>
              <div className={classes.row}>
                  <div className={classes.bulletEmpty}/>
                  <div className={classes.text}>
                      deliberately misleading anyone as to your identity, impersonating another, or falsely identifying the source of any 	information.
                  </div>
              </div>
              <div className={classes.textHeading}>30-day money back guarantee</div>
              <div className={classes.text}>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in.</div>
              <div className={classes.textHeading}>Intellectual Property and Content Ownership</div>
              <div className={classes.text}>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.<br/><br/>Amco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum ip</div>
              <div className={classes.textHeading}>Account Access</div>
              <div className={classes.text}>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.<br/><br/>Amco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum ip</div>
              <div className={classes.textHeading}>Limitation of Liability</div>
              <div className={classNames(classes.text, classes.marginBottom)}>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.<br/><br/>Amco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum ip</div>
          </div>
        </PageMargin>
    </div>
  );
}

export default PrivacyPolicy;