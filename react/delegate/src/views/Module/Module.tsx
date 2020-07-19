import React from 'react';

import { Theme } from 'helpers/theme';
import { createUseStyles, useTheme } from 'react-jss';
import { useRouter } from 'found';
import ModuleMp3 from 'components/Misc/ModuleMp3';
import Page from 'components/Page';

type Props = {
  data: any;
};

const useStyles = createUseStyles((theme: Theme) => ({
  moduleRoot: {
    width: '100%'
  }
}));

const Module = ({ data }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();

  return (
    <Page noPadding>
      <div className={classes.moduleRoot}>
        <ModuleMp3 name="NAME" subTitle="SUBTITle" mp3Url="" />
      </div>
    </Page>
  );
};

export default Module;
