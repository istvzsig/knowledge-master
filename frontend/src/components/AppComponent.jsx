import useFAQs from "../hooks/useFAQs";

import FAQComponent from "./faq/FAQComponent";
import FAQListComponent from "./faq/FAQListComponent";

export default function App() {
  const { faqs, loading, nextFAQ } = useFAQs("faqs");
  return (
    <div>
      {loading ? (
        <h1>Loading...</h1>
      ) : (
        <div>
          <FAQListComponent faqs={faqs} />
          {/* <FAQComponent faq={nextFAQ} /> */}
        </div>
      )}
    </div>
  );
}
