import { useState, useEffect } from "react";
import useSessionStorage from "../hooks/useSessionStorage";

function setFAQCreationDate(faq) {
  faq.createdAt = new Date(faq.createdAt * 1000).toLocaleString();
}

function useFAQs(url) {
  const [error, setError] = useState(undefined);
  const { loadFromSessionStorage, saveToSessionStorage } =
    useSessionStorage(url);
  const [faqs, setFaqs] = useState(loadFromSessionStorage() || []);

  let loading = false;
  let currentFAQIndex = 0;

  async function fetchFAQs() {
    try {
      const response = await fetch("/" + url, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();

      data.forEach((d) => setFAQCreationDate(d));
      setFaqs(data);
      saveToSessionStorage(data);
    } catch (err) {
      console.error("Error fetching FAQs:", err);
      setError(err);
    }
  }

  function getNextFAQ() {
    if (faqs.length > 0) {
      currentFAQIndex = Math.floor((currentFAQIndex + 1) % faqs.length);
    }
  }

  async function loadFAQs() {
    if (faqs.length === 0) {
      loading = true;
      await fetchFAQs();
      loading = false;
    }
  }

  useEffect(() => {
    loadFAQs();
  }, []);

  const nextFAQ = faqs[currentFAQIndex];

  return { faqs, error, loading, nextFAQ, getNextFAQ };
}

export default useFAQs;
