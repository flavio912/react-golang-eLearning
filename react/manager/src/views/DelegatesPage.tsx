import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import Button from 'sharedComponents/core/Input/Button';
import UserSearch from 'components/UserSearch';
import theme, { Theme } from 'helpers/theme';
import PageHeader from 'components/PageHeader';
import Table from 'sharedComponents/core/Table';
import UserLabel from 'sharedComponents/core/Table/UserLabel';
import Text from 'sharedComponents/core/Table/Text/Text';
import Icon from 'sharedComponents/core/Icon';
import CourseCompletion from 'sharedComponents/core/CourseCompletion';
import Dropdown from 'sharedComponents/core/Input/Dropdown';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import Paginator from 'sharedComponents/Pagination/Paginator';
import { createFragmentContainer, graphql } from 'react-relay';
import { Router } from 'found';
import { DelegatesPage_delegates } from './__generated__/DelegatesPage_delegates.graphql';
import { DelegatesPage_manager } from './__generated__/DelegatesPage_manager.graphql';
import { fetchQuery } from 'relay-runtime';
import environment from 'api/environment';
import { OrgOverviewDelegatesQueryResponse } from './__generated__/OrgOverviewDelegatesQuery.graphql';
import { DelegatesPageQueryResponse } from './__generated__/DelegatesPageQuery.graphql';
import { CourseStatus } from './__generated__/DelegateProfilePage_delegate.graphql';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    flexGrow: 1,
    maxWidth: 1275
  },
  searchAndFilter: {
    display: 'flex',
    justifyContent: 'space-between',
    height: 40
  },
  tableOptions: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'
  },
  dropdown: {
    background: 'white'
  },
  divider: {
    width: theme.spacing(1)
  },
  search: {
    flex: 0.4
  },
  actionsRow: {
    borderLeft: '1px solid #ededed',
    paddingLeft: 23,
    height: 38,
    display: 'flex',
    /* justify-content: center; */
    alignItems: 'center'
  }
}));

const delegateRow = (
  userUUID: string,
  name: string,
  profileUrl: string | undefined,
  email: string,
  coursesCompleted: number,
  totalCourses: number,
  lastActiveTimestamp: string,
  nextExpiryTimestamp: string,
  classes: any,
  router: any
): any => ({
  key: userUUID,
  cells: [
    { component: () => <UserLabel name={name} profileUrl={profileUrl} /> },
    {
      component: () => <Text text={email} color={theme.colors.textBlue} />
    },
    {
      component: () => (
        <CourseCompletion total={totalCourses} complete={coursesCompleted} />
      )
    },
    { component: () => <Text text={lastActiveTimestamp} formatDate /> },
    { component: () => <Text text={nextExpiryTimestamp} formatDate /> },
    {
      component: () => (
        <div className={classes.actionsRow}>
          <Icon name={'Card_SecondaryActon_Dots'} />
        </div>
      )
    }
  ],
  onClick: () => router.push(`/app/delegates/${userUUID}`)
});

type Props = {
  delegates: DelegatesPage_delegates;
  manager: DelegatesPage_manager;
  router: Router;
};

const DelegatesPage = ({ delegates, manager, router }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const edges = delegates.edges ?? [];
  const pageInfo = {
    total: delegates.pageInfo?.total ?? 0,
    limit: delegates.pageInfo?.limit ?? 10,
    offset: delegates.pageInfo?.offset ?? 0,
  };
  const page = {
    currentPage: Math.ceil(pageInfo.offset/ pageInfo.limit),
    numPages: Math.ceil(pageInfo.total/ pageInfo.limit)
  };

  const delegateComponents = edges.map((delegate: any) =>
    delegateRow(
      delegate?.uuid,
      `${delegate?.firstName} ${delegate?.lastName}`,
      '',
      delegate?.email,
      (delegate.myCourses as Array<{status: CourseStatus}>).filter(course => course.status == "complete").length,
      delegate.myCourses.length,
      delegate?.lastLogin,
      '2013-04-20T20:00:00+0800',
      classes,
      router
    )
  );
  
  const onUpdatePage = (page: number) => {
    router.push(`/app/delegates?offset=${(page - 1) * pageInfo.limit}&limit=${pageInfo.limit}`);
  };

  return (
    <div className={classes.root}>
      <PageHeader
        showCreateButtons
        title={manager.company.name}
        subTitle="Organisation Overview"
        sideText={`${pageInfo?.total} members total`}
      />
      <div className={classes.searchAndFilter}>
        <div className={classes.search}>
          <UserSearch
            companyName={manager.company.name}
            searchFunction={async (text: string) => {
              const query = graphql`
                query DelegatesPageQuery($name: String!) {
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
              )) as DelegatesPageQueryResponse;

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
        <div className={classes.tableOptions}>
          <Button
            archetype={'default'}
            icon={{ right: 'FilterAdjust' }}
            children={'Filters'}
          />
          <Spacer spacing={1} horizontal />
          <Dropdown
            placeholder={'Sort By'}
            options={[]}
            setSelected={(selected) => <div />}
            className={classes.dropdown}
          />
          <Spacer spacing={1} horizontal />
          <Button
            archetype={'default'}
            icon={{ right: 'DownloadCSV' }}
            noIconPadding
          />
        </div>
      </div>
      <Spacer spacing={3} vertical />
      <Table
        header={[
          'Name',
          'Email',
          'Courses Completed',
          'Last Active',
          'Next Expiry',
          'Actions'
        ]}
        rows={delegateComponents}
      />
      <Spacer vertical spacing={3} />
      <Paginator
        currentPage={page.currentPage}
        updatePage={onUpdatePage}
        numPages={page.numPages}
        itemsPerPage={pageInfo.limit}
        showRange={page.numPages > 4 ? 4 : page.numPages}
      />
    </div>
  );
};

const DelegatesPageFrag = createFragmentContainer(DelegatesPage, {
  delegates: graphql`
    fragment DelegatesPage_delegates on DelegatePage {
      edges {
        uuid
        email
        firstName
        lastName
        lastLogin
        createdAt
        myCourses {
          status
        }
      }
      pageInfo {
        total
        offset
        limit
        given
      }
    }
  `,
  manager: graphql`
    fragment DelegatesPage_manager on Manager {
      company {
        name
      }
    }
  `
});

export default DelegatesPageFrag;
