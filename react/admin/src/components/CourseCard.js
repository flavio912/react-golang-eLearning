import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import {
  Button,
  Card,
  CardContent,
  CardHeader,
  Divider,
  Grid,
  Link,
  Typography,
  colors
} from '@material-ui/core';

const useStyles = makeStyles(theme => ({
  root: {},
  header: {
    paddingBottom: 0
  },
  content: {
    padding: 0,
    '&:last-child': {
      paddingBottom: 0
    }
  },
  description: {
    padding: theme.spacing(2, 3, 1, 3)
  },
  descriptionText: {
    whiteSpace: 'nowrap',
    textOverflow: 'ellipsis',
    overflow: 'hidden'
  },
  tags: {
    padding: theme.spacing(0, 3, 2, 3),
    '& > * + *': {
      marginLeft: theme.spacing(1)
    }
  },
  learnMoreButton: {
    marginLeft: theme.spacing(2)
  },
  likedButton: {
    color: colors.red[600]
  },
  shareButton: {
    marginLeft: theme.spacing(1)
  },
  details: {
    padding: theme.spacing(2, 3)
  },
  capitalise: {
    textTransform: 'uppercase'
  }
}));

function CourseCard({ course, className, ...rest }) {
  const classes = useStyles();

  return (
    <Card {...rest} className={clsx(classes.root, className)}>
      <CardHeader
        className={classes.header}
        disableTypography
        // subheader={
        //   // <Typography variant="body2">
        //   //   by {course.creator.firstName} {course.creator.lastName} | Created:{' '}
        //   //   {moment(course.createdAt).fromNow()}
        //   // </Typography>
        // }
        title={
          <Link
            color="textPrimary"
            component={RouterLink}
            to={`/course/${course.id}/overview`}
            variant="h5"
          >
            {course.name}
          </Link>
        }
      />
      <CardContent className={classes.content}>
        <div className={classes.description}>
          <Typography className={classes.descriptionText}>
            {course.excerpt}
          </Typography>
        </div>
        <div className={classes.tags}>
          {/* {course.tags.map(tag => (
            <Label color={tag.color} key={tag.text}>
              {tag.text}
            </Label>
          ))} */}
        </div>
        <Divider />
        <div className={classes.details}>
          <Grid
            alignItems="center"
            container
            justify="space-between"
            spacing={3}
          >
            <Grid item>
              <Typography variant="h5">£{course.price}</Typography>
              <Typography gutterBottom variant="overline">
                PRICE
              </Typography>
            </Grid>
            <Grid item>
              <div>
                <Typography variant="button" className={classes.capitalise}>
                  {course.type}
                </Typography>
              </div>
              <Typography gutterBottom variant="overline">
                TYPE
              </Typography>
            </Grid>
            <Grid item>
              <Button
                color="primary"
                component={RouterLink}
                size="small"
                to={`/course/${course.id}/overview`}
                variant="outlined"
              >
                View
              </Button>
            </Grid>
          </Grid>
        </div>
      </CardContent>
    </Card>
  );
}

CourseCard.propTypes = {
  className: PropTypes.string,
  course: PropTypes.object.isRequired
};

export default CourseCard;
