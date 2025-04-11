export default function FAQComponent({ faq }) {
  return (
    faq && (
      <div className="faq">
        <p className="faq-id">ID: {faq.id}</p>
        <h3 className="faq-question">{faq.question}?</h3>
        <h3 className="faq-answer">{faq.answer}</h3>
        <p className="faq-created-at">{faq.createdAt}</p>
      </div>
    )
  );
}
