"use client";

import { DesktopNavigation } from "@/components/DesktopNavigation";
import { MobileNavigation } from "@/components/MobileNavigation";

export function Navigation() {
  return (
    <>
      <div className="hidden sm:block">
        <DesktopNavigation />
      </div>
      <div className="block sm:hidden">
        <MobileNavigation />
      </div>
    </>
  );
}
