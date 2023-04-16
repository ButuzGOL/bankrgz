import ky from "ky";
import { useState } from "react";

export default function BankBranchInfo() {
  const [loading, setLoading] = useState<boolean>(false);
  const [data, setData] = useState<any>();
  const [dataReviews, setDataReviews] = useState<any>();
  const [value, setValue] = useState<string>("");

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setValue(e.target.value);
  };

  const onSubmit = async () => {
    setLoading(true);
    try {
      const res = await ky(`bankBranches/${value}`, {
        method: "GET",
        prefixUrl: "http://localhost:7001",
      }).json();

      const resReviews = await ky(`bankBranchReviews/${value}`, {
        method: "GET",
        prefixUrl: "http://localhost:5001",
      }).json();

      setLoading(false);

      // @ts-ignore
      setData(res.Data.data);

      // @ts-ignore
      setDataReviews(resReviews.data);
    } catch (err) {
      // @ts-ignore
      if (err.response?.status === 404) {
        setData(null);
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="relative isolate px-6 pt-14 lg:px-8">
      <div
        className="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
        aria-hidden="true"
      >
        <div
          className="relative left-[calc(50%-11rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
          style={{
            clipPath:
              "polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)",
          }}
        />
      </div>
      <form
        action="#"
        method="POST"
        className="mx-auto mt-16 max-w-xl sm:mt-20 pb-2"
      >
        <div className="grid">
          <div>
            <label
              htmlFor="bank-branch-id"
              className="block text-sm font-semibold leading-6 text-gray-900"
            >
              Bank Branch Id
            </label>
            <div className="mt-2.5">
              <input
                type="text"
                name="bank-branch-id"
                id="bank-branch-id"
                className="block w-full rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                onChange={onChange}
                value={value}
              />
            </div>
          </div>
        </div>
        <div className="mt-10">
          <button
            disabled={loading || !value}
            type="button"
            onClick={onSubmit}
            className="text-white w-full bg-gradient-to-br from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-5 py-3 text-center mr-2 mb-2"
          >
            {loading ? (
              <svg
                aria-hidden="true"
                role="status"
                className="inline w-4 h-4 mr-3 text-white animate-spin"
                viewBox="0 0 100 101"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                  fill="#E5E7EB"
                />
                <path
                  d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                  fill="currentColor"
                />
              </svg>
            ) : null}
            {loading ? "Loading..." : "Submit"}
          </button>
        </div>
      </form>

      {data ? (
        <div className="pb-10 pt-10">
          <a
            href="#"
            className="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700 mx-auto"
          >
            <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              {data.address}
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              N {data.number} <br />
              Dist {data.district} <br />
              Phone {data.phone}
            </p>
          </a>

          <div className="w-full max-w-md p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-8 dark:bg-gray-800 dark:border-gray-700 mx-auto mt-10">
            <div className="flex items-center justify-between mb-4">
              <h5 className="text-xl font-bold leading-none text-gray-900 dark:text-white">
                Reviews
              </h5>
            </div>
            <div className="flow-root">
              {dataReviews.length ? (
                <ul
                  role="list"
                  className="divide-y divide-gray-200 dark:divide-gray-700"
                >
                  {dataReviews.map((item: any) => {
                    return (
                      <li key={item._id} className="py-3 sm:py-4">
                        <div className="flex items-center space-x-4">
                          <div className="flex-shrink-0"></div>
                          <div className="flex-1 min-w-0">
                            <p className="text-sm font-medium text-gray-900 truncate dark:text-white">
                              {item.comment}
                            </p>
                            <p className="text-sm text-gray-500 truncate dark:text-gray-400">
                              {item.rating}
                            </p>
                          </div>
                        </div>
                      </li>
                    );
                  })}
                </ul>
              ) : (
                <p className="text-sm font-medium text-gray-900 truncate dark:text-white">
                  No Reviews
                </p>
              )}
            </div>
          </div>
        </div>
      ) : null}
      {data === null && !loading ? (
        <div className="pb-10 pt-10">
          <a
            href="#"
            className="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700 mx-auto"
          >
            Not found
          </a>
        </div>
      ) : null}
    </div>
  );
}
