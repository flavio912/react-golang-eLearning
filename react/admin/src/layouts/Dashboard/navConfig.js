/* eslint-disable react/no-multi-comp */
/* eslint-disable react/display-name */
import BarChartIcon from '@material-ui/icons/BarChart';
import DashboardIcon from '@material-ui/icons/DashboardOutlined';
import SettingsIcon from '@material-ui/icons/SettingsOutlined';
import LibraryBooks from '@material-ui/icons/LibraryBooks';
import Class from '@material-ui/icons/Class';
import AssignmentInd from '@material-ui/icons/AssignmentInd';
import Layers from '@material-ui/icons/Layers';

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
      }
    ]
  },
  {
    subheader: 'Courses',
    items: [
      {
        title: 'Online Courses',
        href: '/courses',
        icon: LibraryBooks,
        items: [
          {
            title: 'Courses',
            href: '/courses'
          },
          {
            title: 'Modules',
            href: '/courses'
          },
          {
            title: 'Lessons',
            href: '/courses'
          }
        ]
      },
      {
        title: 'Classroom Courses',
        href: '/classroom-courses',
        icon: Class
      },
      {
        title: 'Tutors',
        href: '/tutors',
        icon: AssignmentInd
      }
    ]
  },
  {
    subheader: 'Page Settings',
    items: [
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