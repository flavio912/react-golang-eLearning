import * as React from 'react';
import { createUseStyles } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon';

const useStyles = createUseStyles((theme: Theme) => ({
  tabRoot: {
    position: 'absolute',
    top: '75px',
  },
  offset: {
    top: '-75px',
    left: '265px',
    paddingLeft: '10px',
    zIndex: -1,
  },
  tabDropdown: {
    position: 'relative',
    width: '300px',
    border: ['0.5px', 'solid', theme.colors.borderGrey],
    borderRadius: '8px',
    backgroundColor: theme.colors.primaryWhite,
    paddingTop: '5px',
    paddingbottom: '5px',
    boxShadow: '0px 3px 10px #0000001f',
  },
  tabOption: {
    margin: '20px',
  },
  selectedOption: {
    opacity: 0.5,
  },
  tabTitle: {
    fontSize: theme.fontSizes.large,
    fontWeight: '500',
    marginBottom: '5px',
  },
  tabText: {
    fontSize: theme.fontSizes.small,
    fontWeight: '400',
    color: theme.colors.textGrey,
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
    justifyContent: 'flex-start',
  },
  title: {
    fontSize: theme.fontSizes.large,
    fontWeight: 300,
  },
}));

export interface Tab {
  id: number | string;
  title: string;
  text?: string;
  link?: string;
  options?: Tab[];
}

type Props = {
  tab: Tab;
  selected?: Tab;
  setSelected: (tab?: Tab) => void;
  onClick: (link: string) => void;
};

function TabOption({ tab, selected, setSelected, onClick }: Props) {
  const classes = useStyles();

  const OptionDropdown = (options: Tab[], offset?: boolean) => (
    <div className={classes.tabRoot}>
      <div
        className={classNames(classes.tabDropdown, offset && classes.offset)}
      >
        {options && options.map((option) => Option(option))}
      </div>
    </div>
  );

  const Option = ({ id, title, text, link, options }: Tab) =>
    React.createElement(() => {
      const [selected, setSelected] = React.useState(false);
      return (
        <div
          key={id}
          className={classes.tabOption}
          onClick={() => {
            link && onClick(link);
            options && setSelected(!selected);
          }}
        >
          <div
            className={classNames(
              classes.tabTitle,
              selected && classes.selectedOption,
            )}
          >
            {title}
          </div>
          <div
            className={classNames(
              classes.tabText,
              selected && classes.selectedOption,
            )}
          >
            {text}
          </div>
          {selected && options && OptionDropdown(options, true)}
        </div>
      );
    });

  return (
    <div key={tab.id} className={classes.tab}>
      <div
        className={classes.title}
        onClick={() => {
          if (tab.options) {
            selected && selected.id === tab.id
              ? setSelected(undefined)
              : setSelected(tab);
          } else if (tab.link) {
            onClick(tab.link);
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
      {selected &&
        tab.id === selected.id &&
        tab.options &&
        OptionDropdown(tab.options)}
    </div>
  );
}

export default TabOption;
