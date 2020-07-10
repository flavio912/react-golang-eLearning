import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import FancyBorder from '../../components/Certificate/FancyBorder';
import CertInfo from '../../components/Certificate/CertInfo';
import Signature from '../../components/Certificate/Signature';
import { ReactComponent as Logo } from '../../assets/logo/ttc-logo.svg';
import CertLogoImg from '../../assets/large-1200px-Department_for_Transport.svg.png';

const useStyles = createUseStyles((theme: Theme) => ({
  certGeneratorRoot: {
    margin: 0
  },
  firstPage: {
    marginBottom: 50
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
  className?: string;
  certLogo: string;
  certTitle: string;
  certHead: string;
  certSubHead: string;
  instructorName: string;
  cin: string;
  instructorSign: string;
  certInfo: {
    certName: string,
    moduleDeliver: string,
    forEu: string,
    certNo: string,
    trainingDate: string,
    expiryDate: string
  },
  trainingCompany: {
    title: string,
    website: string,
    email: string,
    phone: string
  }
};

function CertGenerator({ 
  className, 
  certLogo,
  certTitle="UK National Aviation Security Training Programme",
  certHead="Montgomery Demetrius Lucus-Boux",
  certSubHead="Acme Aviation Corporation Limited",
  instructorName="Michelle Waddilove",
  cin="00045618",
  instructorSign="",
  certInfo={
    certName: "General Security Awareness Training (GSAT)",
    moduleDeliver: "Module Delivered: 1-20",
    forEu: "For EU 1998/2015: 11.2.2",
    certNo: "000000054321",
    trainingDate: "DD MMM YYYY",
    expiryDate: "DD MMM YYYY"
  },
  trainingCompany={
    title: "The Training and Compliance Hub Limited",
    website: "www.ttc-hub.com",
    email: "admin@ttc-hub.com",
    phone: "+44 (0)20 3488 2703"
  }
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.certGeneratorRoot}>
      <div className={classes.firstPage}>
        <FancyBorder paperSize="A4">
          <div className={classes.logoArea}>
            <img src={CertLogoImg} className={classes.certLogo} />
            <Logo className={classes.logo} />          
          </div>
          <div className={classes.certHeader}>
            <h1>{certTitle}</h1>
            <p>This is to certify that</p>
            <h3>{certHead}</h3>
            <p>of</p>
            <h4>{certSubHead}</h4>
          </div>
          <CertInfo 
            certName={certInfo.certName}
            moduleDeliver={certInfo.moduleDeliver}
            forEu={certInfo.forEu}
            certNo={certInfo.certNo}
            trainingDate={certInfo.trainingDate}
            expiryDate={certInfo.expiryDate}
          />
          <div className={classes.certInfoFirst}>
            <div className={classes.certRowLeft}>
              <span className={classes.certLabel}>Instructor's Name:</span>
            </div>
            <div className={classes.certRowRight}>
              <span>{instructorName}</span>
              <div><span className={classes.certLabel}>CIN:</span><span>{cin}</span></div>
            </div>
          </div>
          <div className={classes.certRow}>
            <div className={classes.certRowLeft}>
              <span className={classes.certLabel}>Instructor's Signature: </span>
            </div>
            <Signature width={ 400 } height={ 50 } imgSrc={instructorSign} />
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
            <p>For confirmation of the content of this document please contact the training provider.</p>
            <p>For verification of the training provider, please contact the <b>Civil Aviation Authority</b> at avsec.training@caa.gsi.gov.uk</p>
          </div>
        </FancyBorder>
      </div>
      <div>
        <FancyBorder paperSize="A4">
          <div className={classes.logoArea}>
            <img src={CertLogoImg} className={classes.certLogo} />
            <Logo className={classes.logo} />          
          </div>
          <div className={classes.certContent}>
            <div className={classes.leftContent}>
              <p>A) Gov Certificate Logo</p>
              <p>B) Certificate Title</p>
              <p>C) Delegate Name</p>
              <p>D) Company Name</p>
              <p>E) Course Completed (Course Name)</p>
              <p>F) Modules Delivered (count modules assigned to Course)</p>
              <p>G) EU Subtitle Text (added to Course Builder Page at Course Level)</p>
              <p>H) Certificate No. (lifted from a bank of pre-loaded certificate numbers in Admin area, as per Misc.1)</p>
              <p>I) Date of Training (the date they passed)</p>
              <p>J) Expiry Date (at Course level) 13 months, 24 months, 2, 3, 5 years, end of month. Tickbox for end of month.</p>
              <p>K) Instructor/Tutor at Course Level</p>
              <p>L) CIN - certified instructors number, from Tutors table</p>
              <p>M) Signature / tutor level</p>
              <p>N) hard coded TTC address/contact details</p>
            </div>
            <div className={classes.rightContent}>
              <p>Misc. Concerns</p>
              <p>1) Certificate Number <br /> This will be taken from Certificate Regulator Numbers table, 3rd party pre-purchased ext. regular numbers. <span>Number | Type | CertificateID | Used Y/N</span></p>
              <p>2) If doing a joint course, do they get a certificate for the online & the classroom too? Or just 1?</p>
              <p>3) NXCT certificates, page 2 that manually uploaded, is this item per delegate or per Course? Not in spec.</p>
              <p>4) Colour of border for different certs? is this now required? Also, left and right logos - will these changes?</p>
              <p>5) Classroom, student reg.number? Where is this being added or taken from?</p>
            </div>
          </div>
        </FancyBorder>
      </div>
    </div>
  )
}

export default CertGenerator;
