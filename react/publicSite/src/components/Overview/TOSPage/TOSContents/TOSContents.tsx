import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { useRouter } from 'found';
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'flex-start'
    },
    heading: {
        fontSize: theme.fontSizes.large,
        fontWeight: '800',
        marginBottom: '20px',
    },
    link: {
        cursor: 'pointer',
        fontSize: theme.fontSizes.large,
        fontWeight: '600',
        color: theme.colors.textGrey,
        marginBottom: '20px',
    },
    selected: {
        color: theme.colors.navyBlue,
    }
}));

export type LinkDetails = {
    title: string;
    link: string;
}

type Props = {
    links: LinkDetails[];
    selected: LinkDetails;
    setSelected: (link: LinkDetails) => void;
    className?: string;
};

function TOSContents({ links, selected, setSelected, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    const { router } = useRouter();
    const onClick = (link: LinkDetails) => {
        router.push(link.link);
        setSelected(link);
    }

  return (
      <div className={classNames(classes.root, className)}>
          <div className={classes.heading}>CONTENTS</div>
          {links && links.map((link: LinkDetails) => (
              <div
                className={classNames(
                    classes.link,
                    selected === link && classes.selected
                    )}
                onClick={() => onClick(link)}
            >
                {link.title}
            </div>
          ))}
      </div>
  );
}

export default TOSContents;