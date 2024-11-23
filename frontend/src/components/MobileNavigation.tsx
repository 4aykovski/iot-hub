import { MenuIcon } from "lucide-react";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import {
  Drawer,
  DrawerContent,
  DrawerDescription,
  DrawerHeader,
  DrawerTitle,
  DrawerTrigger,
} from "./ui/drawer";
import Link from "next/link";

export function MobileNavigation() {
  return (
    <Drawer direction="left">
      <DrawerTrigger>
        <MenuIcon />
      </DrawerTrigger>
      <DrawerContent>
        <VisuallyHidden>
          <DrawerHeader>
            <DrawerTitle>Navigation Menu</DrawerTitle>
            <DrawerDescription>List with two elements.</DrawerDescription>
          </DrawerHeader>
        </VisuallyHidden>

        <ol className="mt-4 w-full flex flex-col gap-2 justify-center text-center">
          <li>
            <Link href="/" legacyBehavior passHref>
              <h2 className="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0">
                Iot Hub
              </h2>
            </Link>
          </li>
          <li>
            <Link href="/devices" legacyBehavior passHref>
              <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
                Devices
              </h4>
            </Link>
          </li>
        </ol>
      </DrawerContent>
    </Drawer>
  );
}
