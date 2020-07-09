import * as React from "react";
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { useRouter } from 'found';
import { Theme } from 'helpers/theme';
import ArticleHeader from "components/Overview/Article/ArticleHeader";
import ArticleCard, { ArticleDetails } from "components/Overview/Article/ArticleCard/ArticleCard";
import Button from "sharedComponents/core/Input/Button";
import GenreHeader from "components/Overview/Article/GenreHeader";
import PageMargin from "components/core/PageMargin";

const useStyles = createUseStyles((theme: Theme) => ({
    landingRoot: {
        width: '100%',
        backgroundColor: theme.colors.primaryWhite
    },
    header: {
        marginTop: '60px'
    },
    articleGrid: {
        margin: '30px',
        display: 'grid',
        gridTemplateColumns: '1fr 1fr 1fr',
        gridGap: '30px',
        '@media (min-width: 900px) and (max-width: 1200px)': {
            gridTemplateColumns: '1fr 1fr',
        },
        '@media (max-width: 900px)': {
            gridTemplateColumns: '400px',
        }
    },
    button: {
        height: '52px',
        width: '182px',
        margin: '30px 0 60px 0',
        boxShadow: '0 1px 4px 0 rgba(0,0,0,0.09)',
        fontSize: theme.fontSizes.large,
        fontWeight: '800'
    }
}));

const defaultArticle: ArticleDetails = {
    type: "Dangerous Goods",
    date: "20th March 2020",
    description: "CANSO makes call to action to ensure the stability of the ATM industry",
    imageURL: require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")
};

type Props = {};

function ArticleLandingPage(props: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });
  const { router } = useRouter();

  const articles: ArticleDetails[] = [1,2,3,4,5,6,7,8,9].map((item) => ({
    ...defaultArticle,
  }));

  const genres: string[] = ['Latest Articles', 'Popular', 'Dangerous Goods', 'Aviation', 'Health & Safety', 'Product']
  const [selected, setSelected] = React.useState(genres[0]);

  return (
      <div className={classes.landingRoot}>
          <GenreHeader
            genres={genres}
            selected={selected}
            setSelected={setSelected}
        />
          <ArticleHeader
            className={classes.header}
            title="ICAO calls for improved coordination and awareness of COVID-19"
            date="18th April 2020"
            genre="Health & Safety"
            featured="Featured Article"
            image={require("assets/articleHeader.svg")}
          />
          <PageMargin>
            <div className={classes.articleGrid}>
            {articles.map((article: ArticleDetails) => (
                <ArticleCard
                    article={article}
                    onClick={() => router.push('/article')}
                />
            ))}
            </div>
            <Button
            className={classes.button}
            small
            >
                More Articles
            </Button>
        </PageMargin>
      </div>
  );
}

export default ArticleLandingPage;