import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import {
  Modal,
  Card,
  CardHeader,
  CardContent,
  CardActions,
  Grid,
  Divider,
  Typography,
  TextField,
  Button
} from '@material-ui/core';
import { ColorPicker } from 'material-ui-color';
import gql from 'graphql-tag';
import { useMutation, useQuery } from '@apollo/react-hooks';
import ErrorModal from 'src/components/ErrorModal';

const useStyles = makeStyles(theme => ({
  root: {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    outline: 'none',
    boxShadow: theme.shadows[20],
    width: 700,
    maxHeight: '100%',
    overflowY: 'auto',
    maxWidth: '100%'
  },
  actions: {
    justifyContent: 'space-between'
  }
}));

const CREATE_CATEGORY = gql`
  mutation CreateCategory($name: String!, $color: String!) {
    createCategory(input: { name: $name, color: $color }) {
      uuid
    }
  }
`;

const UPDATE_CATEGORY = gql`
  mutation UpdateCategory($uuid: UUID!, $color: String, $name: String) {
    updateCategory(input: { uuid: $uuid, name: $name, color: $color }) {
      uuid
    }
  }
`;

const DELETE_CATEGORY = gql`
  mutation DeleteCategory($uuid: UUID!) {
    deleteCategory(input: { uuid: $uuid })
  }
`;

const GET_CATEGORY = gql`
  query GetCategory($uuid: UUID!) {
    category(uuid: $uuid) {
      uuid
      name
      color
    }
  }
`;

const initialState = {
  name: '',
  color: ''
};

function CategorySaveModal({
  categoryUUID,
  open,
  onClose,
  onSave,
  className,
  ...rest
}) {
  const classes = useStyles();

  const { data, error: fetchErr, refetch, loading } = useQuery(GET_CATEGORY, {
    variables: {
      uuid: categoryUUID
    },
    skip: !categoryUUID,
    fetchPolicy: 'cache-and-network'
  });

  const [formState, setFormState] = useState(
    categoryUUID && !loading ? data.category : initialState
  );
  const [createCategory, { error: createErr }] = useMutation(CREATE_CATEGORY);
  const [updateCategory, { error: updateErr }] = useMutation(UPDATE_CATEGORY);
  const [deleteCategory, { error: deleteErr }] = useMutation(DELETE_CATEGORY);

  useEffect(() => {
    if (!open) return;

    setFormState(categoryUUID && !loading ? data.category : initialState);
  }, [open, categoryUUID, loading]);

  const handleCreateCategory = async event => {
    event.preventDefault();
    try {
      await createCategory({
        variables: {
          name: formState.name,
          color: formState.color
        }
      });
      onSave();
    } catch (err) {}
  };

  const handleUpdateCategory = async event => {
    event.preventDefault();
    try {
      await updateCategory({
        variables: {
          uuid: categoryUUID,
          name: formState.name,
          color: formState.color
        }
      });
      onSave();
    } catch (err) {}
  };

  const handleDeleteCategory = async event => {
    if (
      !window.confirm(
        `Are you sure you want to delete the '${formState.name}' category?`
      )
    )
      return;
    try {
      await deleteCategory({
        variables: {
          uuid: categoryUUID
        }
      });
      onSave();
    } catch (err) {}
  };

  if (loading) {
    return <div></div>;
  }

  return (
    <Modal onClose={onClose} open={open}>
      <Card {...rest} className={clsx(classes.root, className)}>
        <ErrorModal error={fetchErr || createErr || updateErr || deleteErr} />
        <form>
          <CardHeader title={`${categoryUUID ? 'Edit' : 'Create'} Category`} />
          <Divider />
          <CardContent>
            <Grid container spacing={3} direction={'column'}>
              <Grid item>
                <TextField
                  label="Category Name"
                  fullWidth
                  name="name"
                  onChange={evt =>
                    setFormState({ ...formState, name: evt.target.value })
                  }
                  value={formState.name}
                  variant="outlined"
                />
              </Grid>
              <Grid item>
                <Typography variant="overline">Category Color</Typography>
                <ColorPicker
                  value={formState.color}
                  hideTextfield
                  onChange={val =>
                    setFormState({ ...formState, color: `#${val.hex}` })
                  }
                />
              </Grid>
            </Grid>
          </CardContent>
          <Divider />
          <CardActions className={classes.actions}>
            {categoryUUID && (
              <div>
                <Button onClick={handleDeleteCategory} variant="outlined">
                  Delete
                </Button>
              </div>
            )}

            <div className={classes.actions}>
              <Grid container spacing={2}>
                <Grid item>
                  <Button onClick={onClose}>Close</Button>
                </Grid>
                <Grid item>
                  <Button
                    color="primary"
                    onClick={
                      categoryUUID ? handleUpdateCategory : handleCreateCategory
                    }
                    variant="contained"
                  >
                    Save
                  </Button>
                </Grid>
              </Grid>
            </div>
          </CardActions>
        </form>
      </Card>
    </Modal>
  );
}

CategorySaveModal.propTypes = {
  className: PropTypes.string,
  onClose: PropTypes.func,
  open: PropTypes.bool
};

CategorySaveModal.defaultProps = {
  open: false,
  onClose: () => {}
};

export default CategorySaveModal;
