import { useState, useContext, useRef, useEffect } from "react";
import Context from "../../store/context";

const MainSearchBar = () => {
  const [enteredQuery, setEnteredQuery] = useState("");
  const [toggleClearInput, setToggleClearInput] = useState("");
  const inputRef = useRef();

  const { fetchedItems, setSavedItems, savedItems } = useContext(Context);

  const queryInputHandler = (event) => {
    event.preventDefault();
    setEnteredQuery(event.target.value);
  };

  const parser = (query, items) => {
    let filter = [];
    if (query.startsWith("r/")) {
      let parsedQuery = query.replace("r/", "");
      filter = items.filter((item) => {
        return item.subreddit.toLowerCase().includes(parsedQuery);
      });
    } else if (query.startsWith("t/")) {
      let parsedQuery = query.replace("t/", "");
      filter = items.filter((item) => {
        return item.title.toLowerCase().includes(parsedQuery);
      });
    } else if (query.startsWith("u/")) {
      let parsedQuery = query.replace("u/", "");
      filter = items.filter((item) => {
        return item.author.toLowerCase().includes(parsedQuery);
      });
    } else {
      filter = items.filter((item) => {
        return item.selftext.toLowerCase().includes(query);
      });
    }

    return filter;
  };

  useEffect(() => {
    if (enteredQuery === "") {
      setToggleClearInput(false);
    } else {
      setToggleClearInput(true);
    }
  }, [enteredQuery]);

  const filterHandler = (e) => {
    const qlc = enteredQuery.toLowerCase();
    if (e.keyCode === 13) {
      let qarr = qlc.split(";", 2);
      if (qlc.startsWith("n:")) {
        let removed_prefix_qlc = qlc.replace("n:", "");
        setSavedItems(parser(removed_prefix_qlc, savedItems));
      } else {
        if (qarr.length > 1) {
          setSavedItems(parser(qarr[1], parser(qarr[0], fetchedItems)));
        } else {
          setSavedItems(parser(qlc, fetchedItems));
        }
      }
    }
  };

  return (
    <div className="sticky top-4 z-[10] m-auto ">
      <div className="bg-white w-6/12 m-auto grid grid-cols-[repeat(13,_minmax(0,_1fr))] border-[3px] border-black">
        <input
          type="text"
          ref={inputRef}
          onChange={queryInputHandler}
          onKeyDown={filterHandler}
          placeholder={"search " + fetchedItems.length + " items..."}
          className="px-4 py-2 font-mono col-span-12 focus:outline-none"
        />
        <button
          className={toggleClearInput ? "block" : "hidden"}
          onClick={() => {
            setEnteredQuery("");
            inputRef.current.value = "";
            setSavedItems(fetchedItems);
          }}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="m-auto w-7 h-7 "
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M12 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2M3 12l6.414 6.414a2 2 0 001.414.586H19a2 2 0 002-2V7a2 2 0 00-2-2h-8.172a2 2 0 00-1.414.586L3 12z"
            />
          </svg>
        </button>
      </div>
    </div>
  );
};

export default MainSearchBar;
