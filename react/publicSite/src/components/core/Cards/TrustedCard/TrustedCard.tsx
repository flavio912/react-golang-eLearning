import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Card, { PaddingOptions } from 'sharedComponents/core/Cards/Card';
import Icon from 'sharedComponents/core/Icon';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    width: '100%',
    borderRadius: 0
  },
  text: {
    fontSize: theme.fontSizes.extraLargeHeading,
    fontWeight: '800',
    marginBottom: '40px',
    textAlign: 'center'
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    flexWrap: 'wrap'
  },
  noShadow: {
    boxShadow: 'none'
  }
}));

type Props = {
  text?: string;
  noShadow?: boolean;
  padding?: PaddingOptions;
  className?: string;
};

function TrustedCard({
  text,
  noShadow = false,
  padding = 'large',
  className
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <Card
      padding={padding}
      className={classNames(
        classes.root,
        noShadow && classes.noShadow,
        className
      )}
    >
      <div className={classes.text}>{text}</div>
      <div className={classes.row}>
        <Icon
          name="Dhl_Grey"
          style={{ height: '20px', width: '114px', margin: '0 60px 20px 0' }}
        />
        <Icon
          name="Ups_Grey"
          style={{ height: '48px', width: '41px', margin: '0 60px 20px 0' }}
        />
        <Icon
          name="Maersk_Grey"
          style={{ height: '28px', width: '122px', margin: '0 60px 20px 0' }}
        />
        <Icon
          name="DB_Schenker"
          style={{ height: '24px', width: '141px', margin: '0 60px 20px 0' }}
        />
        <Icon
          name="Heathrow"
          style={{ height: '21px', width: '111px', margin: '0 60px 20px 0' }}
        />
        <Icon
          name="Nippon_Express"
          style={{ height: '20px', width: '150px', marginBottom: '20px' }}
        />
      </div>
    </Card>
  );
}

export default TrustedCard;
