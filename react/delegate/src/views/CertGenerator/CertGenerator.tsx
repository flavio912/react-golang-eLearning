import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import FancyBorder from '../../components/Certificate/FancyBorder';
import CertInfo from '../../components/Certificate/CertInfo';
import Signature from '../../components/Certificate/Signature';
import { createFragmentContainer, graphql } from 'react-relay';
import { ReactComponent as Logo } from '../../assets/logo/ttc-logo.svg';
import CertLogoImg from '../../assets/large-1200px-Department_for_Transport.svg.png';
import { CertGenerator_certificateInfo } from './__generated__/CertGenerator_certificateInfo.graphql';

const useStyles = createUseStyles((theme: Theme) => ({
  certGeneratorRoot: {
    margin: 0
  },
  logoArea: {
    height: 100,
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center'
  },
  certLogo: {
    height: 100
  },
  logo: {
    height: 70
  },
  certHeader: {
    marginTop: 20,
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',

    '& h1': {
      fontSize: '2rem',
      fontWeight: 900,
      textAlign: 'center',
      width: 600
    },

    '& p': {
      fontSize: '1.5rem',
      margin: 0
    },

    '& h3': {
      fontSize: '2rem',
      marginTop: '1rem',
      marginBottom: '0'
    },

    '& h4': {
      fontSize: '1.8rem',
      marginTop: 0,
      marginBottom: '1rem'
    }
  },
  certInfoFirst: {
    marginTop: 20,
    display: 'flex'
  },
  certRow: {
    display: 'flex',
    alignItems: 'center',
    marginTop: 15
  },
  certLabel: {
    fontWeight: 900,
    marginRight: 15,
    fontSize: '1.3rem'
  },
  certRowLeft: {
    width: 240,
    textAlign: 'right',
    marginRight: 15
  },
  certRowRight: {
    display: 'flex',
    justifyContent: 'space-between',
    width: 400,

    '& span': {
      fontSize: '1.3rem'
    }
  },
  trainingCompany: {
    display: 'flex',
    marginTop: 15,
    flex: 1
  },
  trainingInfo: {
    width: 400,
    height: 130,
    border: `1px solid ${theme.colors.borderGrey}`,
    padding: 15,
    display: 'flex',
    flexDirection: 'column',
    borderRadius: 5,

    '& h4': {
      fontSize: '1rem',
      marginTop: 0,
      marginBottom: '0.5rem'
    },

    '& span': {
      fontSize: '1.2rem'
    }
  },
  bottomRow: {
    display: 'flex',
    flexDirection: 'column',
    textAlign: 'center',
    margin: '10px 0',

    '& p': {
      margin: 0,
      fontSize: '0.7rem'
    }
  },
  certContent: {
    display: 'flex',
    marginTop: 30,
    marginBottom: 150,
    fontWeight: 'bold'
  },
  leftContent: {
    flex: 1,
    padding: '0 50px 0 20px',
    fontSize: '0.8rem'
  },
  rightContent: {
    flex: 1,
    padding: '0 50px 0 20px',
    fontSize: '0.8rem',

    '& span': {
      fontWeight: 100
    }
  }
}));

type Props = {
  certificateInfo: CertGenerator_certificateInfo;
};

function CertGenerator({ certificateInfo }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const certTitle = 'UK National Aviation Security Training Programme';
  const certHead = 'Montgomery Demetrius Lucus-Boux';
  const certSubHead = 'Acme Aviation Corporation Limited';
  const instructorName = 'Michelle Waddilove';
  const cin = '00045618';
  const instructorSign = '';
  const certData = {
    certName: 'General Security Awareness Training (GSAT)',
    moduleDeliver: 'Module Delivered: 1-20',
    forEu: 'For EU 1998/2015: 11.2.2',
    certNo: '000000054321',
    trainingDate: 'DD MMM YYYY',
    expiryDate: 'DD MMM YYYY'
  };
  const trainingCompany = {
    title: 'The Training and Compliance Hub Limited',
    website: 'www.ttc-hub.com',
    email: 'admin@ttc-hub.com',
    phone: '+44 (0)20 3488 2703'
  };

  return (
    <div className={classes.certGeneratorRoot}>
      <div>
        <FancyBorder paperSize="A4">
          <div className={classes.logoArea}>
            <img src={CertLogoImg} className={classes.certLogo} />
            <Logo className={classes.logo} />
          </div>
          <div className={classes.certHeader}>
            <h1>{certificateInfo.courseTitle}</h1>
            <p>This is to certify that</p>
            <h3>{`${certificateInfo.takerFirstName} ${certificateInfo.takerLastName}`}</h3>
            {certificateInfo.companyName && (
              <>
                <p>of</p>
                <h4>{certificateInfo.companyName}</h4>
              </>
            )}
          </div>
          <CertInfo
            certName={certificateInfo.courseTitle}
            moduleDeliver={certData.moduleDeliver}
            forEu={certData.forEu}
            certNo={certData.certNo}
            trainingDate={certData.trainingDate}
            expiryDate={certData.expiryDate}
          />
          <div className={classes.certInfoFirst}>
            <div className={classes.certRowLeft}>
              <span className={classes.certLabel}>Instructor's Name:</span>
            </div>
            <div className={classes.certRowRight}>
              <span>{instructorName}</span>
              <div>
                <span className={classes.certLabel}>CIN:</span>
                <span>{cin}</span>
              </div>
            </div>
          </div>
          <div className={classes.certRow}>
            <div className={classes.certRowLeft}>
              <span className={classes.certLabel}>
                Instructor's Signature:{' '}
              </span>
            </div>
            <Signature width={400} height={50} imgSrc={instructorSign} />
          </div>
          <div className={classes.trainingCompany}>
            <div className={classes.certRowLeft}>
              <span className={classes.certLabel}>Training Company:</span>
            </div>
            <div className={classes.trainingInfo}>
              <h4>{trainingCompany.title}</h4>
              <span>{trainingCompany.website}</span>
              <span>{trainingCompany.email}</span>
              <span>{trainingCompany.phone}</span>
            </div>
          </div>
          <div className={classes.bottomRow}>
            <p>
              For confirmation of the content of this document please contact
              the training provider.
            </p>
            <p>
              For verification of the training provider, please contact the{' '}
              <b>Civil Aviation Authority</b> at avsec.training@caa.gsi.gov.uk
            </p>
          </div>
        </FancyBorder>
      </div>
    </div>
  );
}

export default createFragmentContainer(CertGenerator, {
  certificateInfo: graphql`
    fragment CertGenerator_certificateInfo on CertificateInfo {
      courseTitle
      expiryDate
      completionDate
      companyName
      takerFirstName
      takerLastName
      certificateBodyURL
      regulationText
      CAANo
      title
      instructorName
      instructorCIN
      instructorSignatureURL
    }
  `
});
