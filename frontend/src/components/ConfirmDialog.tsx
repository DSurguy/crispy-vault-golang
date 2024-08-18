import { Dialog, DialogBackdrop, DialogPanel, DialogTitle } from "@headlessui/react";
import { useState } from "react";

type ConfirmDialogProps = {
  title: React.ReactNode;
  children: React.ReactNode | React.ReactNode[];
  confirmButtonContent?: React.ReactNode;
  denyButtonContent?: React.ReactNode;
  open: boolean;
  onConfirm: () => void;
  onDeny: () => void;
}

export function ConfirmDialog({
  children,
  confirmButtonContent = "Confirm",
  denyButtonContent = "Cancel",
  open,
  onConfirm,
  onDeny,
  title
}: ConfirmDialogProps) {
  const [interactionComplete, setInteractionComplete] = useState(false);

  const handleConfirmClick = () => {
    setInteractionComplete(true);
    onConfirm();
  }

  const handleDenyClick = () => {
    setInteractionComplete(true);
    onDeny();
  }

  return <Dialog open={open} onClose={onDeny} className="relative z-50">
    <DialogBackdrop className="fixed inset-0 bg-black/30" />
    <div className="fixed inset-0 flex w-screen items-center justify-center p-4">
      <DialogPanel className="max-w-lg space-y-4 border bg-white p-12">
        <DialogTitle className="font-bold">{title}</DialogTitle>
        <div>
          {children}
        </div>
        <div className="flex gap-4">
          <button disabled={interactionComplete} onClick={handleConfirmClick}>{confirmButtonContent}</button>
          <button disabled={interactionComplete} onClick={handleDenyClick}>{denyButtonContent}</button>
        </div>
      </DialogPanel>
    </div>
  </Dialog>
}