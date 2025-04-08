import useFAQs from "../../hooks/useFAQs";

export default function FAQComponent({ faq }) {
  return (
    faq && (
      <div className="faq">
        <p className="faq-id">ID: {faq.id}</p>
        <h3 className="faq-question">Question: {faq.question}?</h3>
        <h4 className="faq-answer">Answer: {faq.answer}</h4>
        <h4 className="faq-created-at">Created At: {faq.createdAt}</h4>
      </div>
    )
  );
}
