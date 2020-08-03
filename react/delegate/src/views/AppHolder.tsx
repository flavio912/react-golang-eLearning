import * as React from 'react';
import HeaderMenu from 'components/Menu/HeaderMenu';
import SideMenu from 'components/Menu/SideMenu';
import { Tab } from 'components/Menu/SideMenu/SideMenu';
import { createUseStyles, useTheme } from 'react-jss';
import { useRouter } from 'found';
import SearchResults from 'components/Search/SearchResults';
import { Theme } from 'helpers/theme';
import { createFragmentContainer, graphql, fetchQuery } from 'react-relay';
import { AppHolder_user } from './__generated__/AppHolder_user.graphql';
import environment from 'api/environment';
import { AppHolderQueryResponse } from './__generated__/AppHolderQuery.graphql';
import { SideModalProvider, useSideModalDispatch } from './SideModalProvider';

const useStyles = createUseStyles((theme: Theme) => ({
  appHolder: {
    display: 'flex',
    justifyContent: 'center',
    position: 'relative'
  },
  appHolderRoot: {
    display: 'grid',
    height: '100vh',
    gridTemplateColumns: 'auto 1fr',
    gridTemplateRows: '82px auto'
  },
  appHolderSearch: {
    position: 'absolute',
    width: '100%',
    height: '100%',
    top: 0,
    background: theme.searchBackground
  }
}));

type Props = {
  children?: React.ReactChildren;
  user?: AppHolder_user;
};

const AppHolder = ({ children, user }: Props) => {
  const classes = useStyles();
  const { match, router } = useRouter();
  const tabs: Tab[] = [
    { id: 0, icon: 'LeftNav_Icon_Dashboard', title: 'Dashboard' },
    { id: 1, icon: 'LeftNav_Icon_Courses', title: 'Online Courses', size: 23 },
    {
      id: 2,
      icon: 'LeftNav_Icon_Training',
      title: 'Training Progress',
      size: 23
    }
  ];

  const selected = () => {
    const { routes } = match;
    const currentRouter = routes[routes.length - 1];
    switch (currentRouter.path) {
      case '/':
        return tabs[0];
      case '/courses':
        return tabs[1];
      case '/courses/:id':
        return tabs[1];
      case '/progress':
        return tabs[2];
      default:
        return tabs[0];
    }
  };
  const [isShowSearchModal, setIsShowSearchModal] = React.useState<boolean>(
    false
  );
  const onToggleSearchModal = () => setIsShowSearchModal(!isShowSearchModal);
  const hideSearch = () => setIsShowSearchModal(false);

  const dispatchModal = useSideModalDispatch();
  return (
    <div className={classes.appHolderRoot}>
      <SideMenu
        tabs={tabs}
        selected={selected()}
        logo={require('../assets/logo/ttc-logo.svg')}
        onClick={(tab) => {
          switch (tab.id) {
            case 0:
              router.push('/app');
              break;
            case 1:
              router.push('/app/courses');
              break;
            case 2:
              router.push('/app/progress');
            default:
              break;
          }
        }}
      />
      <HeaderMenu
        user={{ name: `${user?.firstName} ${user?.lastName}`, url: '' }}
        onToggleSearchModal={onToggleSearchModal}
        showSearch={user?.type === 'individual'}
      />
      <div className={classes.appHolder}>
        {children}
        {isShowSearchModal && (
          <div className={classes.appHolderSearch} onClick={hideSearch}>
            <SearchResults
              searchFunction={async (text: string, offset: number) => {
                const query = graphql`
                  query AppHolderQuery($name: String!, $offset: Int!) {
                    courses(
                      filter: { name: $name }
                      page: { limit: 4, offset: $offset }
                    ) {
                      edges {
                        ident: id
                        name
                        bannerImageURL
                        introduction
                        price
                      }
                      pageInfo {
                        total
                        limit
                        offset
                        given
                      }
                    }
                  }
                `;

                const variables = {
                  name: text,
                  offset: offset
                };

                const data = (await fetchQuery(
                  environment,
                  query,
                  variables
                )) as AppHolderQueryResponse;
                if (
                  !data ||
                  !data.courses ||
                  !data.courses.edges ||
                  !data.courses.pageInfo
                ) {
                  console.error('Could not get data', data);
                  return {
                    resultItems: [],
                    pageInfo: {
                      currentPage: 1,
                      numPages: 1
                    }
                  };
                }

                const resultItems = data.courses.edges.map((course) => ({
                  id: course?.ident ?? '',
                  title: course?.name ?? '',
                  image: course?.bannerImageURL ?? '',
                  description: course?.introduction ?? '',
                  onClick: () => {
                    if (
                      course?.ident === undefined ||
                      course?.name === undefined ||
                      course?.price === undefined
                    ) {
                      console.error('Unable to get course', course);
                      return;
                    }

                    dispatchModal({
                      type: 'show',
                      courses: [
                        {
                          id: course.ident,
                          name: course.name,
                          price: course.price
                        }
                      ]
                    });
                  }
                }));

                const pageInfo = {
                  currentPage: Math.ceil(data.courses.pageInfo.offset / 4 + 1),
                  numPages: Math.ceil(data.courses.pageInfo.total / 4)
                };
                const result = {
                  resultItems: resultItems,
                  pageInfo: pageInfo
                };

                return result;
              }}
            />
          </div>
        )}
      </div>
    </div>
  );
};

export default createFragmentContainer(AppHolder, {
  user: graphql`
    fragment AppHolder_user on User {
      type
      firstName
      lastName
    }
  `
});
