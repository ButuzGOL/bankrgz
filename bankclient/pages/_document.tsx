import { Html, Head, Main, NextScript } from "next/document";

const navigation = [
  { name: "Bank Branch Info", href: "/bankBranchInfo" },
  { name: "Bank Branch Delete", href: "/bankBranchDelete" },
];

export default function Document() {
  return (
    <Html lang="en">
      <Head>
        <title>Bank Branches</title>
      </Head>
      <body>
        <div className="bg-white">
          <header className="absolute inset-x-0 top-0 z-50">
            <nav
              className="flex items-center justify-between p-6 lg:px-8"
              aria-label="Global"
            >
              <div className="flex lg:flex-1">
                <a href="#" className="-m-1.5 p-1.5">
                  <span className="sr-only">Your Company</span>
                  <img
                    className="h-8 w-auto"
                    src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
                    alt=""
                  />
                </a>
              </div>

              <div className="hidden lg:flex lg:gap-x-12">
                {navigation.map((item) => (
                  <a
                    key={item.name}
                    href={item.href}
                    className="text-sm font-semibold leading-6 text-gray-900"
                  >
                    {item.name}
                  </a>
                ))}
              </div>
              <div className="hidden lg:flex lg:flex-1 lg:justify-end"></div>
            </nav>
          </header>
          <Main />
        </div>
        <NextScript />
      </body>
    </Html>
  );
}
