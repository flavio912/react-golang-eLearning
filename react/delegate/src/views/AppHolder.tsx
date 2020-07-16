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
import { AppHolderQuery, AppHolderQueryResponse } from './__generated__/AppHolderQuery.graphql';

const useStyles = createUseStyles((theme: Theme) => ({
  appHolder: {
    display: 'flex',
    padding: '42px 60px',
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

const results = [
  {
    id: 1,
    title: 'Cargo Manager (CM) – VC, HS, XRY, EDS',
    image: 'https://www.gstatic.com/webp/gallery/1.jpg',
    description:
      'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
  },
  {
    id: 2,
    title: 'Cargo Manager (CM) – VC, HS, XRY, EDS',
    image: 'https://www.gstatic.com/webp/gallery/1.jpg',
    description:
      'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
  },
  {
    id: 3,
    title: 'Cargo Manager (CM) – VC, HS, XRY, EDS',
    image: 'https://www.gstatic.com/webp/gallery/1.jpg',
    description:
      'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
  }
];

type Props = {
  children?: React.ReactChildren;
  user?: AppHolder_user;
};

const AppHolder = ({ children, user }: Props) => {
  console.log('USER', user);
  const classes = useStyles();
  const { match, router } = useRouter();
  const tabs: Tab[] = [
    { id: 0, icon: 'LeftNav_Icon_Dashboard', title: 'Dashboard' },
    { id: 1, icon: 'LeftNav_Icon_Courses', title: 'Online Courses', size: 23 },
    {
      id: 2,
      icon: 'LeftNav_Icon_Courses',
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
      />
      <div className={classes.appHolder}>
        {children}
        {isShowSearchModal && (
          <div className={classes.appHolderSearch} onClick={hideSearch}>
            <SearchResults searchFunction={
              async (text: string) => {
                const query = graphql`
                  query AppHolderQuery($name: String!){
                    courses(filter: { name: $name }, page: {limit: 4}){
                      edges{
                        notId: id
                        name
                        bannerImageURL
                        introduction
                      }
                      pageInfo{
                        total
                        limit
                        offset
                        given
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
                )) as AppHolderQueryResponse;
                
                if (!data || !data.courses || !data.courses.edges || !data.courses.pageInfo){
                  console.error('Could not get data', data);
                  return {
                    resultItems: [],
                    pageInfo: {
                      totalPages: 1,
                      offset: 0,
                      limit: 4,
                      totalItems: 0
                    }
                  };
                }

                const resultItems = data.courses.edges.map((course) => ({
                  id: course?.notId ?? '',
                  title: course?.name ?? '',
                  image: course?.bannerImageURL ?? '',
                  description: course?.introduction ?? ''
                }));
                
                const pageInfo = {
                  totalPages: data.courses.pageInfo?.total,
                  offset: data.courses.pageInfo.offset,
                  limit: data.courses.pageInfo.limit,
                  totalItems: data.courses.pageInfo.given
                };

                const result = {
                  resultItems: resultItems,
                  pageInfo: pageInfo
                }

                return result;
              }
            } />
          </div>
        )}
      </div>
    </div>
  );
};

export default createFragmentContainer(AppHolder, {
  user: graphql`
    fragment AppHolder_user on User {
      firstName
      lastName
    }
  `,
  courses: graphql`
    fragment AppHolder_courses on CoursePage {
      edges {
        notId: id
        name
        bannerImageURL
        introduction
      }
      pageInfo {
        total
        offset
        limit
        given
      }
    }
  `,
});
