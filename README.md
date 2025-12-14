# conf2
Enhancing distributed systems with multi-version concurrency control over 2PL protocol

# storage_optimization
**STORAGE OPTIMIZATION IN DISTRIBUTED ENVIRONMENTS USING OPTIMISTIC CONCURRENCY CONTROL**
* Author: Vipul Reddy
* Published In : International Journal on Science and Technology (IJSAT)
* Publication Date: 06-2024
* E-ISSN: 2229-7677
* Impact Factor: 9.88
* Link:

**Abstract:**\
This paper addresses performance and storage overhead challenges in database transaction management caused by maintaining multiple data versions under Multi-Version Concurrency Control (MVCC). It examines how the accumulation of record versions and associated garbage collection processes introduce storage overhead and management complexity, particularly in high-concurrency environments. The study emphasizes the trade-offs of MVCC, including version proliferation, cleanup costs, and the impact of complex transaction interactions such as write skew. By integrating Optimistic Concurrency Control, the proposed approach reduces unnecessary version creation and minimizes storage overhead while preserving snapshot isolation and consistency guarantees. The paper highlights the need for efficient version management strategies to enhance scalability, throughput, and resource utilization in high-performance database systems.

**Key Contributions:**
* **Storage Overhead Mitigation:**\
Investigated the storage inefficiencies introduced by Multi-Version Concurrency Control due to version proliferation and garbage collection overhead in distributed environments.

* **Optimistic Concurrency Integration:**\
Applied Optimistic Concurrency Control selectively to reduce unnecessary version creation while preserving transactional consistency and isolation guarantees.
  
* **Comprehensive Evaluation:** \
  Conducted a detailed comparative analysis of MVCC and OCC across multiple cluster sizes, demonstrating significant reductions in storage consumption and predictable scalability behavior.
  
* **Research & Implementation Leadership:**\
  Directed the design, implementation, and experimental validation of a hybrid concurrency control approach focused on improving storage efficiency and scalability.

**Relevance & Real-World Impact**
* **Improved Storage Efficiency:**\
Reduced disk space usage in distributed database systems by minimizing versioning overhead without compromising transaction correctness.

* **Scalable Distributed Deployment:**\
Enabled more storage-efficient scaling of database clusters by adopting OCC in low-contention scenarios while retaining MVCC where strong consistency is required.

* **System-Level Performance Benefits:** \
    Lowered garbage collection pressure and metadata management costs, resulting in improved throughput and reduced resource utilization.
  
* **Academic and Educational Value:** \
    Provides empirical insights and implementation references for research and teaching in concurrency control, distributed databases, and storage optimization.

**Experimental Results (Summary)**:

  | Nodes | Multi-Version Concurrency Control Storage | Optimistic Concurrency Control | Reduction (%)   |
  |-------|-------------------------------------------| -------------------------------| ----------------|
  | 3     |  3                                        | 1                              | 66.67           |
  | 5     |  5                                        | 1.5                            | 70.00           |
  | 7     |  7                                        | 2                              | 71.43           |
  | 9     |  9                                        | 2.5                            | 72.22           |
  | 11    |  11                                       | 3                              | 72.73           |

**Citation** \
STORAGE OPTIMIZATION IN DISTRIBUTED ENVIRONMENTS USING OPTIMISTIC CONCURRENCY CONTROL
* Vipul R 
* International Journal on Science and Technology 
* E-ISSN 2229-7677
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijsat.org/ \
**Author Contact** \
**LinkedIn**: http://linkedin.com/in/Please add here | **Email**: please keep email id @gmail.com







