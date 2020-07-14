import React, { useEffect } from 'react';
import {
  Card,
  CardHeader,
  TextField,
  CardContent,
  Button,
  Grid,
  Divider
} from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import { Autocomplete } from '@material-ui/lab';
import TagsInput from 'src/components/TagsInput';
import ErrorModal from 'src/components/ErrorModal';

const GET_CATEGORIES = gql`
  query GetCategories($limit: Int!, $text: String) {
    categories(page: { limit: $limit }, text: $text) {
      edges {
        uuid
        name
        color
      }
    }
  }
`;

function CourseInfo({ state, setState }) {
  let categoryOptions = [{ title: 'Aviation Security', value: 'avsec' }];

  const { loading, error, data, refetch } = useQuery(GET_CATEGORIES, {
    variables: {
      limit: 100
    },
    fetchPolicy: 'cache-and-network'
  });

  if (!loading && !error) {
    categoryOptions = data.categories.edges.map(category => ({
      title: category.name,
      value: category.uuid
    }));
  }

  return (
    <Card>
      <ErrorModal error={error} />
      <CardHeader title={'Course Info'} />
      <Divider />
      <CardContent>
        <Grid container direction="column" spacing={2}>
          <Grid item>
            <TextField
              fullWidth
              label="Course Name"
              name="courseName"
              onChange={inp => {
                setState({ name: inp.target.value });
              }}
              placeholder="e.g Dangerous Goods"
              value={state.name}
              variant="outlined"
            />
          </Grid>
          <Grid item>
            <Grid container spacing={1} alignItems={'center'}>
              <Grid item xs={12}>
                <Autocomplete
                  options={categoryOptions}
                  loading={loading}
                  getOptionLabel={option => option.title}
                  onChange={(event, newValue) => {
                    if (!newValue) return;
                    setState({ categoryUUID: newValue.value });
                  }}
                  renderInput={params => (
                    <TextField
                      {...params}
                      label="Primary Course Category"
                      variant="outlined"
                    />
                  )}
                />
              </Grid>
            </Grid>
          </Grid>
          <Grid item>
            <Grid container spacing={1} alignItems={'center'}>
              <Grid item xs={12}>
                <Autocomplete
                  options={categoryOptions}
                  getOptionLabel={option => option.title}
                  onChange={(event, newValue) => {
                    setState({ secondaryCategory: newValue.value });
                  }}
                  renderInput={params => (
                    <TextField
                      {...params}
                      label="Secondary Course Category"
                      variant="outlined"
                    />
                  )}
                />
              </Grid>
            </Grid>
          </Grid>
          <Grid item>
            <TagsInput onChange={newVal => setState({ tags: newVal })} />
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default CourseInfo;
