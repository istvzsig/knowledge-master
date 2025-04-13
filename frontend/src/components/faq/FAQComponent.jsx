export default function FAQComponent({ faq }) {
  if (!faq && !faq.id) return;

  if (typeof faq.createdAt !== "string") {
    faq.createdAt = new Date(faq.createdAt * 1000).toLocaleString();
  }

  return (
    <div className="faq">
      <p className="faq-id">ID: {faq.id}</p>
      <h3 className="faq-question">{faq.question}?</h3>
      <h3 className="faq-answer">{faq.answer}</h3>
      <p className="faq-created-at">{faq.createdAt}</p>
    </div>
  );
}
