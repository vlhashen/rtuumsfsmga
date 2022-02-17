import { useContext, useState, useRef, useEffect } from "react";
import Context from "../../store/context";

const FilterSearchBar = () => {
  const { setDisplayItem, subredditKeys, clearToggle, setClearSelected } =
    useContext(Context);
  const [toggleClearInput, setToggleClearInput] = useState(false);
  const [keyInput, setKeyInput] = useState("");
  const inputRef = useRef();

  const queryInputHandler = (event) => {
    event.preventDefault();
    setKeyInput(event.target.value);
  };

  useEffect(() => {
    const qlc = keyInput.toLowerCase();
    if (qlc === "") {
      setDisplayItem(subredditKeys);
      setToggleClearInput(false);
    } else {
      setToggleClearInput(true);
      let filter = subredditKeys.filter((item) => {
        return item.toLowerCase().includes(qlc);
      });

      if (filter.length > 0) {
        setDisplayItem(filter);
      } else {
        setDisplayItem([]);
      }
    }
  }, [keyInput, setDisplayItem, subredditKeys]);

  return (
    <div className="flex w-3/5 mt-6 mb-6 ml-4">
      <div className="grid sm:grid-cols-5 lg:grid-cols-8 w-6/12 pl-4 py-2 font-mono tracking-tighter border-[2px] border-black border-dashed">
        <input
          type="text"
          placeholder="filter subreddit"
          onChange={queryInputHandler}
          className="sm:col-span-4 lg:col-span-7 focus:outline-none"
          ref={inputRef}
        />
        <button
          className={toggleClearInput ? "block" : "hidden"}
          onClick={() => {
            setKeyInput("");
            inputRef.current.value = "";
          }}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="w-6 h-6"
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
      <span className={clearToggle ? "block" : "hidden"}>
        <button
          className="flex p-1 ml-2"
          onClick={() => {
            setClearSelected(true);
          }}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="w-8 h-8 fill-red-400 stroke-black"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fillRule="evenodd"
              d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
              clipRule="evenodd"
            />
          </svg>
        </button>
      </span>
    </div>
  );
};

export default FilterSearchBar;
