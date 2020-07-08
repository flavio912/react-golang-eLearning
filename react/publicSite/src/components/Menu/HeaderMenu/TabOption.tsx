import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import Icon from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
    tabRoot: {
        position: 'absolute',
        top: '75px',
        zIndex: 10
    },
    tabDropdown: {
        position: 'relative',
        right: '25px',
        maxWidth: '300px',
        border: ['0.5px', 'solid', theme.colors.borderGrey],
        borderRadius: '8px',
        backgroundColor: theme.colors.primaryWhite,
        padding: '5px 0',
        boxShadow: '0px 3px 10px #0000001f'
    },
    tabOption: {
        margin: '20px'
    },
    tabTitle: {
        fontSize: theme.fontSizes.large,
        fontWeight: '500',
        marginBottom: '5px'
    },
    tabText: {
        fontSize: theme.fontSizes.small,
        fontWeight: '400',
        color: theme.colors.textGrey
    },
    tab: {
        fontFamily: 'Muli',
        fontSize: theme.fontSizes.large,
        fontWeight: '300',
        marginRight: '30px',
        cursor: 'pointer',
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'flex-start'
    },
    title: {
        fontSize: theme.fontSizes.large,
        fontWeight: 300
    },
}));

export interface Tab {
    id: number;
    title: string;
    link?: string;
    options?: {
      title: string;
      text: string;
      link: string;
    }[];
  }

type Props = {
    tab: Tab;
    selected?: Tab;
    setSelected: (tab?: Tab) => void;
    onClick: (link: string) => void;
};

function TabOption({ tab, selected, setSelected, onClick }: Props) {
  const classes = useStyles();
  return (
    <div
        key={tab.id}
        className={classes.tab}
    >
        <div
            className={classes.title}
            onClick={() => {
                if (tab.options) {
                    selected && selected.id === tab.id
                    ? setSelected(undefined)
                    : setSelected(tab);
                } else if (tab.link) {
                    onClick(tab.link)
                }
            }}
        >
            {tab.title}
        </div>
        {tab.options && (
            <Icon
                name="Down_Arrow"
                size={10}
                style={{ cursor: 'pointer', marginLeft: '5px' }}
            />
        )}
        {selected && tab.id === selected.id && tab.options && (
            <div className={classes.tabRoot}>
                <div className={classes.tabDropdown}>
                {tab.options.map(option => (
                    <div
                        className={classes.tabOption}
                        onClick={() => onClick(option.link)}
                    >
                        <div className={classes.tabTitle}>{option.title}</div>
                        <div className={classes.tabText}>{option.text}</div>
                    </div>
                ))}
                </div>
            </div>
        )}
    </div>
  );
}

export default TabOption;