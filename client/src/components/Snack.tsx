import ErrorIcon from '@material-ui/icons/Error';
import React from 'react';
import { Snackbar, SnackbarContent } from '@material-ui/core';


type SnackProps = {
  error: string,
  onClose?: () => void,
}
export function ErrorSnack({ error, onClose }: SnackProps) {

  return (
    <Snackbar open={error ? true : false} autoHideDuration={3000} onClose={onClose}>
      <SnackbarContent
        style={{ backgroundColor: '#d32f2f', padding: 10, color: 'white', display: 'flex' }}
        role="alert"
        message={
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <ErrorIcon style={{ marginRight: 10 }} />
            <span> {error} </span>
          </div>
        }
      />
    </Snackbar>
  )
}