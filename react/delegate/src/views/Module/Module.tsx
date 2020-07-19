import React from 'react';

import { Theme } from 'helpers/theme';
import { createUseStyles, useTheme } from 'react-jss';
import { useRouter } from 'found';
import ModuleMp3 from 'components/Misc/ModuleMp3';
import Page from 'components/Page';
import PageTitle from 'components/PageTitle';
import { createFragmentContainer, graphql } from 'react-relay';
import { Module_myActiveCourse } from './__generated__/Module_myActiveCourse.graphql';
import { Module_module } from './__generated__/Module_module.graphql';
import CourseSyllabusCardFrag from 'components/Overview/CourseSyllabusCard/CourseSyllabusCardFrag';

const useStyles = createUseStyles((theme: Theme) => ({
  moduleRoot: {
    width: '100%'
  },
  mp3: {
    background: 'white',
    height: 111,
    padding: '0px 42px',
    borderBottom: `1px solid ${theme.colors.borderGrey}`
  },
  contentHolder: {
    width: '100%',
    display: 'flex',
    position: 'relative',
    justifyContent: 'center'
  },
  content: {
    display: 'grid',
    gridGap: 50,
    maxWidth: 1040,
    gridTemplateColumns: '2fr 1fr'
  }
}));

type Props = {
  myActiveCourse: Module_myActiveCourse;
  module: Module_module;
};

const Module = ({ myActiveCourse, module }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();
  console.log('MOD', module);
  return (
    <Page noPadding>
      <div className={classes.moduleRoot}>
        <ModuleMp3
          name="NAME"
          subTitle="SUBTITle"
          mp3Url=""
          className={classes.mp3}
        />
        <div className={classes.contentHolder}>
          <div className={classes.content}>
            <PageTitle title={'COOLIO'} />
            {/* <CourseSyllabusCardFrag course={myActiveCourse?.course} /> */}
          </div>
        </div>
      </div>
    </Page>
  );
};

export default createFragmentContainer(Module, {
  myActiveCourse: graphql`
    fragment Module_myActiveCourse on MyCourse {
      course {
        ...CourseSyllabusCardFrag_course
      }
    }
  `,
  module: graphql`
    fragment Module_module on Module {
      name
      uuid
    }
  `
});
