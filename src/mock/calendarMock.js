import uuid from 'uuid/v1';
import moment from 'moment';
import { colors } from '@material-ui/core';
import mock from 'src/utils/mock';

mock.onGet('/api/calendar').reply(200, {
  draft: [],
  events: [
    {
      id: uuid(),
      title: 'Dangerous Goods Course v2',
      desc:
        'Some important information about the course that will be shown to users',
      color: colors.green['700'],
      start: moment('2020-03-22 16:55:00'),
      end: moment('2020-03-22 18:02:00')
    },
    {
      id: uuid(),
      title: 'Handling Course v4',
      desc:
        'Some important information about the course that will be shown to users',
      color: colors.green['700'],
      start: moment('2020-03-24 16:55:00'),
      end: moment('2020-03-24 18:02:00')
    }
  ]
});
