# In-Memory SQL-like Database

## Problem Statement

You have been tasked with designing and implementing an **in-memory SQL-like database** that supports basic table management and querying operations.

---

## Requirements

### Table Management
- **Create a Table**
  - Tables can be dynamically created at runtime.
  - Each table must have:
    - A unique name.
    - One or more column definitions.
---

### Column Definition
- Each column must have:
  - A name.
  - A **data type** (e.g., `INT`, `STRING`, `FLOAT`, `BOOL`).

---

### Data Insertion
- Users can insert rows into tables.
- The system must:
  - Validate the input data against the column data types.

---

### Querying
- Users can:
  - Fetch all rows from a table.
---

## Features Summary

| Feature             | Description                                           |
|---------------------|-------------------------------------------------------|
| Table Creation      | Define table with column names, types, and constraints |
| Insert Rows         | Add new rows to a table with data validation          |
| Type Validation     | Enforce column data types on insert                   |
| Query by Column     | Retrieve rows by column-based filter                  |

---

## Example

### Create Table
```sql
CREATE TABLE users (
  id INT NOT NULL UNIQUE,
  name STRING NOT NULL,
  email STRING UNIQUE
);
