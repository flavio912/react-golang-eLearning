import * as React from "react";
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from "helpers/theme";
import Icon from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
    genreCenterer: {
        position: 'fixed',
        width: '100%',
        display: 'flex',
        justifyContent: 'center',
        backgroundColor: theme.colors.backgroundGrey,
    },
    genreRow: {
        display: 'flex',
        width: theme.centerColumnWidth,
        '@media (max-width: 800px)': {
            display: 'none'
        }
    },
    genre: {
        cursor: 'pointer',
        fontSize: theme.fontSizes.large,
        color: theme.colors.textGrey,
        margin: '20px 40px 20px 0',
        '@media (max-width: 800px)': {
            margin: '20px'
        }
    },
    selected: {
        color: theme.colors.primaryBlack,
        fontWeight: 'bold'
    },
    mobileGenre: {
        display: 'flex',
        width: theme.centerColumnWidth,
        paddingLeft: '25px',
        '@media (min-width: 800px)': {
            display: 'none'
        }
    },
    genreDropdown: {
        position: 'absolute',
        top: '50px',
        left: '25px',
        right: '25px',
        border: ['0.5px', 'solid', theme.colors.borderGrey],
        borderRadius: '8px',
        backgroundColor: theme.colors.backgroundGrey,
        padding: '5px 0',
        boxShadow: '0px 3px 10px #0000001f'
    },
    backgroundHider: {
        position: 'absolute',
        width: '100%',
        height: '50%',
        top: 0,
        left: 0
      }
}));

type Props = {
    genres: string[];
    selected: string;
    setSelected: (genre: string) => void;
    className?: string
};

function GenreHeader({ genres, selected, setSelected, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    const [showPopup, setShowPopup] = React.useState(false);
  return (
    <>
        <div className={classes.backgroundHider} onClick={() => setShowPopup(false)} />
        <div className={classNames(classes.genreCenterer, className)}>
            <div className={classes.genreRow}>
                {genres.map((genre: string) => (
                    <div className={classNames(
                        classes.genre,
                        selected === genre && classes.selected
                    )}
                        onClick={() => setSelected(genre)}
                    >
                        {genre}
                    </div>
                ))}
            </div>
            <div className={classes.mobileGenre}>
                <div
                    className={classNames(classes.genre, classes.selected)}
                    onClick={() => setShowPopup(!showPopup)}
                >
                    {selected}
                    <Icon
                        name="Down_Arrow"
                        size={10}
                        style={{ cursor: 'pointer', marginLeft: '10px' }}
                    />
                </div>
                {showPopup && (
                    <div className={classes.genreDropdown}>
                        {genres.map((genre: string) => (
                            <div
                                className={classes.genre}
                                onClick={() => {
                                    setSelected(genre);
                                    setShowPopup(false);
                                }}
                            >
                                {genre}
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    </>
  );
}

export default GenreHeader;