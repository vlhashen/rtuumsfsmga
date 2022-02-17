import { forwardRef } from "react";
import ReactMarkdown from "react-markdown";
import BookmarkButton from "../buttons/BookmarkButton";

const SavedComment = forwardRef((props, ref) => {
  const {
    title,
    body,
    timestamp,
    link,
    saved,
    subreddit,
    author,
    name,
    concise,
  } = props;
  const parsedDate = new Date(timestamp * 1000).toLocaleString();

  return (
    <div>
      <div
        ref={ref}
        className="flex-1 p-3 font-mono bg-red-200 border-2 border-black border-dashed hover:border-solid hover:border-3"
      >
        <div className="grid grid-cols-3">
          <div className="h-4 text-xs leading-3 col-span-2">
            <p className="mb-[0.075rem] font-bold">r/{subreddit}</p>
            <h1 className="tracking-tight">by u/{author}</h1>
          </div>
          { concise ? (
            <div className="h-7"></div>
          ) : (
            <BookmarkButton saved={saved} name={name} />
          )}
        </div>

        <a href={link} target="_blank" rel="noreferrer">
          <div>
            <h2 className="mr-2 font-bold break-normal leading-5">{title}</h2>
            <p className="text-xs tracking-tighter">{parsedDate}</p>
            <div className="p-2 overflow-hidden text-sm text-center">
              <ReactMarkdown>
                {body.length < 499 ? body : body + "..."}
              </ReactMarkdown>
            </div>
          </div>
        </a>
      </div>
    </div>
  );
});

export default SavedComment;
