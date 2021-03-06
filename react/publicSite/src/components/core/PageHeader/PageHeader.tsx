import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { useRouter } from 'found';
import { Theme } from 'helpers/theme';
import Button from 'sharedComponents/core/Input/Button';
import Icon from 'sharedComponents/core/Icon';
import PageMargin from '../PageMargin';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: theme.colors.lightBlue,
    padding: '57px 0',
  },
  defaultTitle: {
    alignSelf: 'center',
    fontSize: theme.fontSizes.extraLarge,
    fontWeight: '800',
    textAlign: 'center',
    borderBottom: '3px solid #0f62e8',
    paddingBottom: 10,
    marginBottom: 30,
  },
  defaultDesc: {
    fontSize: '40px',
    fontWeight: '800',
    marginTop: '10px',
    maxWidth: '950px',
    textAlign: 'center',
  },
  buttonsTitle: {
    fontSize: '40px',
    border: 'none',
  },
  buttonsDesc: {
    fontSize: theme.fontSizes.xSmallHeading,
    fontWeight: '400',
  },
  bar: {
    alignSelf: 'center',
    width: '55px',
    height: '3px',
    backgroundColor: theme.colors.navyBlue,
  },
  jumpText: {
    margin: '20px 20px 0 20px',
  },
  button: {
    fontWeight: '800',
    margin: '20px 20px 0 20px',
    height: '53px',
    width: '211px',
  },
  buttons: {
    flexWrap: 'wrap',
    '@media (max-width: 500px)': {
      flexDirection: 'column',
    },
  },
  history: {
    alignSelf: 'flex-start',
    marginBottom: '50px',
  },
  extraLarge: {
    fontSize: theme.fontSizes.extraLarge,
  },
  bold: {
    fontWeight: 'bold',
    marginLeft: '3px',
  },
  row: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
  },
  column: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'flex-start',
    alignItems: 'center',
  },
}));

export type Archetypes = 'default' | 'buttons';

export type ButtonLink = {
  title: string;
  link: string;
};

type Props = {
  title: string;
  description: string;
  archetype?: Archetypes;
  history?: string[];
  buttons?: ButtonLink[];
  className?: string;
};

function PageHeader({
  title,
  description,
  archetype,
  history,
  buttons,
  className,
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const titleStyle = classes[archetype + 'Title'];
  const descStyle = classes[archetype + 'Desc'];

  const { router } = useRouter();
  const onClick = (link?: string) => {
    link && router.push(link);
  };

  return (
    <div className={classNames(classes.root, className)}>
      <PageMargin>
        {history && (
          <div className={classNames(classes.row, classes.history)}>
            {history.map((page: string, index: number) =>
              index !== history.length - 1 ? (
                <div className={classes.extraLarge}>
                  {page} <Icon name="Right_Arrow" size={12} />
                </div>
              ) : (
                <div className={classNames(classes.extraLarge, classes.bold)}>
                  {page}
                </div>
              ),
            )}
          </div>
        )}

        <div className={classes.row}>
          <div className={classes.column}>
            <div className={classNames(classes.defaultTitle, titleStyle)}>
              {title}
            </div>
            {archetype && archetype === 'default' && (
              <div className={classes.bar} />
            )}
            <div className={classNames(classes.defaultDesc, descStyle)}>
              {description}
            </div>

            {archetype && archetype === 'buttons' && (
              <div className={classNames(classes.row, classes.buttons)}>
                <div
                  className={classNames(classes.jumpText, classes.extraLarge)}
                >
                  Jump to:
                </div>
                {buttons &&
                  buttons.map((buttonLink: ButtonLink) => (
                    <Button
                      className={classNames(classes.button, classes.extraLarge)}
                      onClick={() => onClick(buttonLink.link)}
                    >
                      {buttonLink.title}
                    </Button>
                  ))}
              </div>
            )}
          </div>
        </div>
      </PageMargin>
    </div>
  );
}

export default PageHeader;
