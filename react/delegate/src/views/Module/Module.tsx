import React, { useState } from 'react';

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
import SelectButton from 'components/core/Input/SelectButton';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import Button from 'components/core/Input/Button';
import { goToNextURL } from 'views/helpers';

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
  },
  nextQuestionWrap: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: 150,
    background: 'white',
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: 6
  }
}));

type Props = {
  myActiveCourse: Module_myActiveCourse;
  module?: Module_module;
};

const Module = ({ myActiveCourse, module }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router, match } = useRouter();
  const { courseID, moduleUUID } = match.params;

  const selectOptions = ['Description', 'Transcript'];
  const [selectedOption, setSelectedOption] = useState(selectOptions[0]);

  return (
    <Page noPadding>
      <div className={classes.moduleRoot}>
        <ModuleMp3
          name={module?.name ?? ''}
          subTitle=""
          mp3Url={module?.voiceoverURL ?? ''}
          className={classes.mp3}
        />
        <Spacer vertical spacing={3} />
        <div className={classes.contentHolder}>
          <div className={classes.content}>
            <div>
              <PageTitle title={module?.name ?? ''} />
              <Spacer vertical spacing={3} />
              <SelectButton
                options={selectOptions}
                selected={selectedOption}
                onClick={(option: string) => {
                  setSelectedOption(option);
                }}
              />
              <Spacer vertical spacing={3} />
              <div>
                {selectedOption == selectOptions[0]
                  ? module?.description
                  : module?.transcript}
              </div>
              <Spacer vertical spacing={3} />
              <div className={classes.nextQuestionWrap}>
                <Button
                  title={'Continue'}
                  padding="large"
                  onClick={() => {
                    router.push(
                      goToNextURL(
                        parseInt(courseID),
                        myActiveCourse.course.syllabus,
                        module?.uuid
                      )
                    );
                  }}
                />
              </div>
            </div>
            <div>
              <CourseSyllabusCardFrag
                course={myActiveCourse?.course}
                upTo={myActiveCourse?.upTo ?? undefined}
              />
            </div>
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
        syllabus {
          name
          type
          uuid
          ... on Module {
            syllabus {
              name
              uuid
              type
            }
          }
        }
        ...CourseSyllabusCardFrag_course
      }
      upTo
    }
  `,
  module: graphql`
    fragment Module_module on Module {
      name
      uuid
      voiceoverURL
      description
      transcript
    }
  `
});
