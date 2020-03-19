/* eslint-disable react/no-multi-comp */
/* eslint-disable react/display-name */
import React from 'react';
import { colors } from '@material-ui/core';
import BarChartIcon from '@material-ui/icons/BarChart';
import CalendarTodayIcon from '@material-ui/icons/CalendarToday';
import DashboardIcon from '@material-ui/icons/DashboardOutlined';
import SettingsIcon from '@material-ui/icons/SettingsOutlined';
import LibraryBooks from '@material-ui/icons/LibraryBooks';
import Class from '@material-ui/icons/Class';
import AssignmentInd from '@material-ui/icons/AssignmentInd';
import Layers from '@material-ui/icons/Layers';
import Label from 'src/components/Label';

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
        title: 'Calendar',
        href: '/calendar',
        icon: CalendarTodayIcon,
        label: () => <Label color={colors.green[500]}>New</Label>
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
