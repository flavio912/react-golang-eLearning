import * as React from 'react';
import ActvityCard from 'components/Overview/ActivityCard';
import { createUseStyles, useTheme } from 'react-jss';
import UserSearch from 'components/UserSearch';
import TitleWrapper from 'components/Overview/TitleWrapper';
import QuickOverview from 'components/Overview/QuickOverview';
import TrainingProgressCard from 'components/Overview/TrainingProgressCard';
import { Theme } from 'helpers/theme';
import ProfileCard from 'components/Overview/ProfileCard';
import PageHeader from 'components/PageHeader';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import { fetchQuery } from 'relay-runtime';
import { createFragmentContainer, graphql } from 'react-relay';
import { OrgOverview_manager } from './__generated__/OrgOverview_manager.graphql';
import environment from 'api/environment';
import { OrgOverviewDelegatesQueryResponse } from './__generated__/OrgOverviewDelegatesQuery.graphql';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    flexGrow: 1,
    maxWidth: 1275
  },
  activity: {},
  statsRow: {
    display: 'flex',
    justifyContent: 'space-between',
    flexWrap: 'wrap'
  },
  cardFlex: {
    display: 'flex',
    justifyContent: 'space-between'
  },
  breaker: {
    width: theme.spacing(2)
  },
  searchRow: {
    display: 'flex',
    flexWrap: 'wrap',
    height: 40
  },
  bottomRow: {
    display: 'flex',
    justifyContent: 'space-between',
    flexWrap: 'wrap'
  },
  spacer: {
    minWidth: 770
  },
  grid: {
    display: 'grid',
    gridTemplateRows: '40px auto auto',
    gridGap: 38,
    gridTemplateAreas: `
      "search search search search search .      .      .      .      .      .     "
      "overvw overvw overvw overvw overvw .      traini traini traini traini traini"
      "profle profle profle profle profle actvty actvty actvty actvty actvty actvty"
    `,
    '@media (max-width: 1350px)': {
      gridTemplateAreas: `
      "search"
      "overvw"
      "traini"
      "profle"
      "actvty"
    `
    }
  },
  search: {
    flex: 1,
    gridArea: 'search'
  },
  quickOverview: {
    gridArea: 'overvw'
  },
  trainingProgress: {
    gridArea: 'traini'
  },
  yourInfo: {
    gridArea: 'profle'
  },
  activityWrapper: {
    gridArea: 'actvty'
  }
}));

type Props = {
  manager: OrgOverview_manager;
};

const OrgOverview = ({ manager }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.root}>
      <PageHeader
        showCreateButtons
        title={manager.company.name}
        subTitle="Organisation Overview"
      />
      <div className={classes.grid}>
        <div className={classes.search}>
          <UserSearch
            companyName={manager.company.name}
            searchFunction={async (text: string) => {
              const query = graphql`
                query OrgOverviewDelegatesQuery($name: String!) {
                  delegates(filter: { name: $name }, page: { limit: 8 }) {
                    edges {
                      uuid
                      TTC_ID
                      firstName
                      lastName
                    }
                  }
                }
              `;

              const variables = {
                name: text
              };

              const data = (await fetchQuery(
                environment,
                query,
                variables
              )) as OrgOverviewDelegatesQueryResponse;

              if (!data || !data.delegates || !data.delegates.edges) {
                console.error('Could not get data', data);
                return [];
              }

              const results = data.delegates?.edges.map((delegate) => ({
                uuid: delegate?.uuid ?? '',
                key: `${delegate?.firstName} ${delegate?.lastName}`,
                value: delegate?.TTC_ID ?? ''
              }));

              return results;
            }}
          />
        </div>
        <TitleWrapper title="Quick Overview" className={classes.quickOverview}>
          <QuickOverview
            purchasedCourses={20}
            numDelegates={manager.company.delegates?.pageInfo?.total ?? 0}
            numValidCertificates={10}
            numCertificatesExpiringSoon={15}
          />
        </TitleWrapper>
        <TitleWrapper
          title="Training Progress"
          className={classes.trainingProgress}
        >
          <div className={classes.cardFlex}>
            <TrainingProgressCard
              coursesDone={20}
              timeTracked={{ h: 20, m: 15 }}
              title="Weekly"
            />
            <Spacer spacing={2} horizontal />
            <TrainingProgressCard
              coursesDone={20}
              timeTracked={{ h: 20, m: 15 }}
              title="Monthly"
            />
          </div>
        </TitleWrapper>
        <TitleWrapper title="Your information" className={classes.yourInfo}>
          <ProfileCard
            heading="Profile"
            fields={[
              {
                fieldName: 'Name',
                value: `${manager.firstName} ${manager.lastName}`
              },
              { fieldName: 'Role', value: 'Group Leader' },
              { fieldName: 'Email', value: manager.email },
              { fieldName: 'Tel Contact', value: manager.telephone },
              {
                fieldName: 'Active since',
                value: new Date(manager.createdAt || '').toDateString()
              }
            ]}
            padding="medium"
          />
        </TitleWrapper>
        <TitleWrapper title={'Activity'} className={classes.activityWrapper}>
          <ActvityCard
            className={classes.activity}
            padding={'medium'}
            leftHeading={'Certificate status'}
            rightHeading={'Recent Updates'}
            options={[]}
            updates={[]}
            data={{
              outerRing: {
                name: 'Current',
                value: 154
              },
              innerRing: {
                name: 'Expired',
                value: 64
              }
            }}
          />
        </TitleWrapper>
      </div>
    </div>
  );
};

export default createFragmentContainer(OrgOverview, {
  manager: graphql`
    fragment OrgOverview_manager on Manager {
      firstName
      lastName
      email
      telephone
      createdAt
      company {
        name
        delegates {
          pageInfo {
            total
          }
        }
      }
    }
  `
});
