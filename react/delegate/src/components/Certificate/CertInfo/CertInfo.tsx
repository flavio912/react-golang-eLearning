import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";

type Props = {
  certName: string,
  moduleDeliver: string,
  forEu: string,
  certNo: string,
  trainingDate: string,
  expiryDate: string
};

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: 'flex',
    boxSizing: 'border-box',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    background: theme.colors.certBackgroundGrey,
    border: `1px solid ${theme.colors.borderGrey}`,
    padding: '15px 0'
  },
  certName: {
    fontSize: '1.3rem',
    margin: '1rem 0',
    color: theme.colors.primaryBlue
  },
  certInfoRow: {
    margin: 0,
    fontSize: '1.3rem',

    '& span': {
      fontSize: '1.1rem',
      float: 'left'
    }
  },
  certInfoGroup: {
    marginBottom: '10px'
  },
  certInfoLabel: {
    width: '150px',
    fontWeight: 900,
    textAlign: 'right',
    display: 'block',
    float: 'left',
    marginRight: '10px',
    fontSize: '1.7rem'
  }
}));

function CertInfo({
  certName,
  moduleDeliver,
  forEu,
  certNo,
  trainingDate,
  expiryDate
}: Props) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      <p className={classes.certInfoRow}>Has successfully completed:</p>
      <h4 className={classes.certName}>{certName}</h4>
      <div className={classes.certInfoGroup}>
        <p className={classes.certInfoRow}>{moduleDeliver}</p>
        <p className={classes.certInfoRow}>{forEu}</p>
      </div>
      <div className={classes.certInfoGroup}>
        <p className={classes.certInfoRow}>
          <span className={classes.certInfoLabel}>Certificate No:</span>
          <span>{certNo}</span>
        </p>
        <p className={classes.certInfoRow}>
          <span className={classes.certInfoLabel}>Date of Training:</span>
          <span>{trainingDate}</span>
        </p>
        <p className={classes.certInfoRow}>
          <span className={classes.certInfoLabel}>Expiry Date:</span>
          <span>{expiryDate}</span>
        </p>
      </div>
    </div>
  );
}

export default CertInfo;
