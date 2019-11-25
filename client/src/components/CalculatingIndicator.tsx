import React from 'react';
import { Dialog, DialogContent, CircularProgress } from '@material-ui/core';

type CompSelectorProps={
  open: boolean,
}
export default function CalculatingIndicator({open}: CompSelectorProps) {
  return (
    <Dialog
      open={open}
    >
      <DialogContent style={{width: 100, height: 100, justifyContent: 'center', alignItems: 'center', display: 'flex'}}>
        <CircularProgress />
      </DialogContent>
    </Dialog>
  )
}