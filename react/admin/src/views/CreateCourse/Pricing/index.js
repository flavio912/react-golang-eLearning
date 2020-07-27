import React from 'react';
import {
  Grid,
  TextField,
  Card,
  CardHeader,
  CardContent,
  Divider,
  InputAdornment,
  Switch,
  Typography
} from '@material-ui/core';
import { Autocomplete } from '@material-ui/lab';
import { makeStyles } from '@material-ui/styles';
import SideOptions from './SideOptions';
import { gql } from 'apollo-boost';
import { useQuery } from '@apollo/react-hooks';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  buttonText: {
    color: '#4a4a4a',
    fontSize: 11,
    fontWeight: 'weight: 700'
  },
  termsInput: {
    width: '100%'
  },
  expirations: {
    alignItems: 'center'
  }
}));

const GET_CERTIFICATE_TYPES = gql`
  query GetCategories($limit: Int!, $name: String) {
    certificateTypes(page: { limit: $limit }, filter: { name: $name }) {
      edges {
        uuid
        name
      }
    }
  }
`;

function Pricing({ state, setState }) {
  const classes = useStyles();

  const { loading, error, data } = useQuery(GET_CERTIFICATE_TYPES, {
    variables: {
      limit: 100
    },
    fetchPolicy: 'cache-and-network'
  });

  let certOptions = [];

  if (!loading && !error) {
    certOptions = data.certificateTypes.edges.map(certificateType => ({
      name: certificateType.name,
      uuid: certificateType.uuid
    }));
  }

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item={6} xs={8}>
          <Grid container spacing={2} direction={'column'}>
            <Grid item>
              <Card>
                <CardHeader title={'Terms and conditions'} />
                <Divider />
                <CardContent>
                  <TextField
                    label=""
                    multiline
                    className={classes.termsInput}
                    rows={5}
                    value={state.terms}
                    onChange={inp => {
                      setState({ terms: inp.target.value });
                    }}
                    placeholder={'Terms and conditions'}
                    variant="outlined"
                  />
                </CardContent>
              </Card>
            </Grid>
            <Grid item>
              <Card>
                <CardHeader title={'Certificate Options'} />
                <Divider />
                <CardContent>
                  <Grid container spacing={4} direction="column">
                    <Grid item>
                      <Autocomplete
                        value={state.certificateType}
                        options={certOptions}
                        loading={loading}
                        getOptionLabel={option => option.name}
                        onChange={(event, newValue) => {
                          console.log('new', newValue);
                          if (!newValue) return;
                          setState({
                            certificateType: {
                              uuid: newValue.uuid,
                              name: newValue.name
                            }
                          });
                        }}
                        renderInput={params => (
                          <TextField
                            {...params}
                            fullWidth
                            label="Certificate Type"
                            variant="outlined"
                          />
                        )}
                      />
                    </Grid>
                    <Grid
                      container
                      item
                      className={classes.expirations}
                      spacing={4}
                    >
                      <Grid item>
                        <TextField
                          label="Expires in"
                          InputProps={{
                            endAdornment: (
                              <InputAdornment position="end">
                                Months
                              </InputAdornment>
                            )
                          }}
                          variant="outlined"
                          type="number"
                          value={state.expiresInMonths}
                          onChange={evt => {
                            try {
                              setState({
                                expiresInMonths: parseFloat(evt.target.value)
                              });
                            } catch (err) {}
                          }}
                        />
                      </Grid>
                      <Grid item>
                        <Typography variant="overline">
                          To end of month
                        </Typography>
                        <Typography>
                          Take the expiration date from the end of the month
                        </Typography>
                        <Switch
                          checked={state.expirationToEndMonth}
                          color="secondary"
                          name="RequiresCAA"
                          onChange={(evt, checked) => {
                            setState({ expirationToEndMonth: checked });
                          }}
                          value={state.expirationToEndMonth}
                        />
                      </Grid>
                    </Grid>
                  </Grid>
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={4}>
          <Grid container direction={'column'} spacing={2}>
            <Grid item>
              <SideOptions state={state} setState={setState} />
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default Pricing;
