import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import React from "react";
import { Label } from "recharts";
import { updateDevice } from "../../../../server/api";

interface UpdateDeviceBtnProps {
  id: string;
}

const UpdateDeviceBtn: React.FC<UpdateDeviceBtnProps> = ({ id }) => {
  const [name, setName] = React.useState("");
  const [limit, setLimit] = React.useState("");

  function handleUpdate() {
    const limitNumber = Number(limit);
    updateDevice(id, { name, limit: limitNumber });
    setName("");
    setLimit("");
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline">Изменить</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Изменение данных устройства</DialogTitle>
          <DialogDescription></DialogDescription>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label name="name" className="text-right">
              Название
            </Label>
            <Input
              id="name"
              defaultValue={name}
              placeholder="Температура кухня"
              className="col-span-3"
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label name="limit" className="text-right">
              Лимит
            </Label>
            <Input
              id="limin"
              defaultValue={limit}
              placeholder="40"
              className="col-span-3"
              onChange={(e) => setLimit(e.target.value)}
            />
          </div>
        </div>
        <DialogFooter>
          <DialogClose asChild>
            <Button type="submit" onClick={handleUpdate}>
              Сохранить изменения
            </Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default UpdateDeviceBtn;
