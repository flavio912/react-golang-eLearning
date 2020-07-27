import * as React from "react";
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from "helpers/theme";
import ArticleHeader from "components/Overview/Article/ArticleHeader";
import GenreHeader from "components/Overview/Article/GenreHeader";
import PageMargin from "components/core/PageMargin";

const useStyles = createUseStyles((theme: Theme) => ({
    articleRoot: {
        width: '100%',
        backgroundColor: theme.colors.primaryWhite
    },
    header: {
        marginTop: '60px'
    },
    text: {
        fontSize: theme.fontSizes.extraLarge,
        color: theme.colors.textGrey,
        margin: '50px 0'
    }
}));

type Props = {};

function Article(props: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });
  const genres: string[] = ['Latest Articles', 'Popular', 'Dangerous Goods', 'Aviation', 'Health & Safety', 'Product']
  const [selected, setSelected] = React.useState(genres[0]);

  // Replace with text from API
  const text = "The International Civil Aviation Organisation (ICAO) has made new calls to governments across the globe, emphasising the need for improved coordination with aircraft operators on the current air services updates and flight restrictions that are in force as a result of the current coronavirus (COVID-19) crisis.Additionally, ICAO has asked its member states to examine the best means of supporting stakeholders from the aviation sector, including maintenance, air traffic services, and other safety- and security- critical aviation system suppliers.Dr. Fang Lui, ICAO Secretary General, stressed: “These are truly unprecedented times and they are posing risks to the airline operator and airport profitability that most passengers would be familiar with. As COVID-19 continues to impede and diminish global mobility in all world regions, we’re also seeing very serious risks emerging to the operational viability of air traffic control systems and safety oversight systems, vital support industry segments such as ground services, repair and maintenance facilities, and other key system providers.”The calls for improved government- operator coordination were featured in ICAO’s most recent state letter. It drew member states’ attention to how some current flight crew notifications issued by states were not providing sufficient detail on the respective national flight operations restrictions, airport closures and reductions in air traffic services that are currently in force.";
  
  return (
    <div className={classes.articleRoot}>
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
            image={require("assets/articleHeader.svg")}
            author={{ name: "James Green", url: require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")}}
        />
        <PageMargin>
            <div className={classes.text}>
                {text}
            </div>
        </PageMargin>
    </div>
  );
}

export default Article;