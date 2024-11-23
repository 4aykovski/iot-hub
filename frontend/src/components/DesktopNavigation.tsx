import Link from "next/link";
import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
} from "./ui/navigation-menu";
import { Button } from "./ui/button";

export function DesktopNavigation() {
  return (
    <NavigationMenu>
      <NavigationMenuList className="gap-6">
        <NavigationMenuItem>
          <Link href="/" legacyBehavior passHref>
            <NavigationMenuLink>
              <Button variant="link">
                <h3 className="scroll-m-20 text-2xl font-semibold tracking-tight">
                  Iot Hub
                </h3>
              </Button>
            </NavigationMenuLink>
          </Link>
        </NavigationMenuItem>
        <NavigationMenuItem>
          <Link href="/devices" legacyBehavior passHref>
            <NavigationMenuLink>
              <Button variant="link">
                <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
                  Devices
                </h4>
              </Button>
            </NavigationMenuLink>
          </Link>
        </NavigationMenuItem>
      </NavigationMenuList>
    </NavigationMenu>
  );
}
