import { ModeToggle } from "@/components/ModeToggle";
import { ThemeProvider } from "./ThemeProvider";
import "./globals.css";
import { Navigation } from "./Navigation";
import { Toaster } from "@/components/ui/sonner";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning={true}>
      <head></head>
      <body>
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <header className="flex px-10 py-5 justify-between">
            <Navigation />
            <ModeToggle />
          </header>
          {children}
          <Toaster />
        </ThemeProvider>
      </body>
    </html>
  );
}
