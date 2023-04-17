# Strategy

## General Strategy

1. Gather information from the user:
   * Understand the user's request: Analyze user input and identify keywords, topics, and requirements.
   * Enrich the information: Ask follow-up questions to clarify the user's needs and preferences.
   * Generate a detailed job description: Create a comprehensive outline of the task, including objectives, criteria, and user preferences.
   * Confirm understanding: Present the job description to the user for verification and modify it if necessary.
   * Compress the final description: Utilize compression/shortening to reduce the token count while retaining essential information.

2. Task segmentation and prioritization:
   * Analyze the given task: Break down the task into smaller, manageable subtasks.
   * Determine dependencies: Identify relationships and dependencies among subtasks.
   * Prioritize subtasks: Rank subtasks based on their importance, urgency, and dependencies.
   * Assign token limits: Allocate a portion of the 4k token limit to each subtask, considering complexity and dependencies.

3. Information retrieval and summarization:
   * Conduct a Google Search: Use the Google Search API to find relevant title/description/URL results for each subtask.
   * Scrape websites: Deploy the scraper to extract data from the identified websites.
   * Index and search with Pinecone: Utilize Pinecone to index and search the scraped data, identifying relevant information for each subtask.
   * Summarize data: Use a secondary AI instance to summarize and compress the data, maintaining token limits for each subtask.

4. Data storage and retrieval optimization:
   * Create a structured filesystem: Design an organized filesystem to store and categorize data for each subtask.
   * Store and read files: Implement filesystem commands to save and access files in the working directory.
   * Implement caching: Introduce a caching mechanism to reduce redundancy and enhance AI performance.

5. Utilizing multiple AI instances for collaboration:
   * Assign roles to AI instances: Designate primary and secondary AI instances to manage different aspects of the task.
   * Primary AI instance: Task segmentation, prioritization, information retrieval, and storage optimization.
   * Secondary AI instance: Data summarization and compression.
   * Establish communication protocol: Develop an efficient method for AI instances to collaborate and share information.

6. Task execution and feedback:
   * Execute subtasks: Complete each prioritized subtask using the retrieved, summarized, and compressed data.
   * Monitor progress: Track the completion status of subtasks and adjust priorities if necessary.
   * Review and feedback: Request user feedback on completed subtasks to ensure satisfaction and make adjustments as needed.

