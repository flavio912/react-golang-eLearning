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
    padding: '20px 0'
  },
  certName: {
    color: theme.colors.primaryBlue
  },
  certInfoRow: {
    margin: 0
  },
  certInfoGroup: {
    marginBottom: '10px'
  },
  certInfoLabel: {
    width: '140px',
    fontWeight: 'bold',
    textAlign: 'right',
    display: 'block',
    float: 'left',
    marginRight: '10px'
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
