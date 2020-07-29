/* eslint-disable react/no-multi-comp */
/* eslint-disable react/display-name */
import BarChartIcon from '@material-ui/icons/BarChart';
import DashboardIcon from '@material-ui/icons/DashboardOutlined';
import SettingsIcon from '@material-ui/icons/SettingsOutlined';
import LibraryBooks from '@material-ui/icons/LibraryBooks';
import Class from '@material-ui/icons/Class';
import AssignmentInd from '@material-ui/icons/AssignmentInd';
import Layers from '@material-ui/icons/Layers';
import CardMembershipIcon from '@material-ui/icons/CardMembership';

export default [
  {
    subheader: 'Management',
    items: [
      {
        title: 'Dashboard',
        href: '/dashboard',
        icon: DashboardIcon
      },
      {
        title: 'Companies',
        href: '/companies',
        icon: BarChartIcon
      },
      {
        title: 'Individuals',
        href: '/individuals',
        icon: AssignmentInd
      },
      {
        title: 'Delegates',
        href: '/delegates',
        icon: AssignmentInd
      },
      {
        title: 'Company requests',
        href: '/approve-companies',
        icon: BarChartIcon
      }
    ]
  },
  {
    subheader: 'Courses',
    items: [
      {
        title: 'Courses',
        href: '/courses',
        icon: LibraryBooks,
        items: [
          {
            title: 'Courses',
            href: '/courses'
          },
          {
            title: 'Modules',
            href: '/modules'
          },
          {
            title: 'Tests',
            href: '/tests'
          },
          {
            title: 'Questions',
            href: '/questions'
          },
          {
            title: 'Lessons',
            href: '/lessons'
          }
        ]
      },
      {
        title: 'Certificate Types',
        href: '/certificateTypes',
        icon: CardMembershipIcon
      },
      {
        title: 'Tutors',
        href: '/tutors',
        icon: AssignmentInd
      }
    ]
  },
  {
    subheader: 'Admin Settings',
    items: [
      {
        title: 'Admins',
        href: '/admins',
        icon: AssignmentInd
      },
      {
        title: 'Page Editor',
        href: '/page-editor',
        icon: Layers
      },
      {
        title: 'Global Settings',
        href: '/settings',
        icon: SettingsIcon
      }
    ]
  }
];
