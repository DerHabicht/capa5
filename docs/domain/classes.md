# Class Diagrams

```mermaid
---
title: Package Hierarchy
---

classDiagram
    coordination ..> user
    plan ..> user
    user ..> unit

    class activity {
    }

    class billet {
    }

    class budget {
    }

    class coordination {

    }

    class plan {
        Plan
        PT
        RevisionHistory
        Revision
        Section
    }

    class task {
        Task
        CoordinationChain
        Coordination
    }

    class unit {
        AB
        DutyAssignment
        Member
        Unit
    }

    class user {
        User
    }

```

## Activity Module

## Billeting Module

## Budgeting Module

## Coordination Module
```mermaid
classDiagram
    StaffSummarySheet o-- Action
    StaffSummarySheet o-- Tab
    Action *-- CA
    Action *-- CO
    Tab o-- Task
    Tab *-- Tabbable
    Tab o-- Comment
    Task -- Action
    Comment *-- CT
    Comment *-- CD

    class StaffSummarySheet {
        -actions: []Action
        -actionOfficer: user.User
        -plan: plan.Plan
        -suspense: time.Time  
        -summary: string
        -tabs: []Tab
        -crm: []Comment
    }

    class Action {
        -to: user.User
        -action: CA
        -outcome: CO
        -date: time.Time
    }

    class CA {
	    ApproveAction
	    CoordAction
	    InfoAction
	    SignAction
	    ReviewAction
	    POCAction
	    LogAction
    }

    class CO {
        ConcurOutcome
        NonConcurOutcome
        ApproveOutcome
        DisapproveOutcome
    }
   
    class Tab {
        item: Tabbable
    }

    class Tabbable {
        ID() uuid.UUID
        TabType() string
    }

    class Task {
        -actions: []Action
    }

    class Comment {
        source: user.User
        opr: user.User
        ct: CT
        tab: Tab
        location: string
        comment: string
        cd: CD
        rationale: string
    }

    class CT {
        CriticalComment
        MajorComment
        SubstantiveComment
        AdministrativeComment
    }

    class CD {
        AcceptDecision
        RejectDecision
        ModifyDecision
    }


```


## Planning Module

```mermaid
classDiagram
    Plan *-- PT
    Plan *-- RevisionHistory
    RevisionHistory o-- Revision
    Revision *-- Section

    class Plan {
        -planNumber: string
        -planType: PT
        -parentPlan: *Plan
        -title: string
        -owner: user.User
        -editors: collections.Set[user.User]
        -revisionHistory: RevisionHistory
    }

    class PT {
        BPLAN
        CONPLAN
        OPLAN
        OPORD
        FRAGORD
    }

    class RevisionHistory {
        -draftRevision: *Revision
        -publishedRevisions: collections.Stack[Revision]
        +Draft(create bool) *Revision
        +PublishDraft() bool
        +Latest() Revision
        +Revision(revisionNumber int) Revision, error
        +RevisionNotes() []map[string]string
    }

    class Revision {
        -revisionNumber: int
        -owner: user.User
        -editors: collections.Set[user.User]
        -effective: *time.Time
        -revisionNotes: string
        -sections: []Section
        -annexes: []Section
        +RevisionNumber() int
        +Effective() *time.Time
        +RevisionNotes() map[string]string
        +Sections() []Section
        +Annexes() []Section
        +AddSection(s Section, index int)
        +AddAnnex(a Section, index int)
        +RemoveSection(index int)
        +RemoveAnnex(index int)
    }

    class Section {
        -title: string
        -owner: user.User
        -editors: collections.Set[user.User]
        -content: string
        -subsections: []Section
        +Title() string
        +Content() string
        +UpdateContent(content string) string, error
        +Subsections() []Section
        +AddSubsection(s Section, index int)
        +RemoveSubsection(index int)
    }
```

## Unit Module

Age Brackets:

- Adult: 21+
- Early Adult: 18-20
- Late Teen: 16-17
- Mid Teen: 14-15
- Early Teen: 12-13

```mermaid
classDiagram
    class AB {
        Adult
        EarlyAdult
        LateTeen
        MidTeen
        EarlyTeen
    }

    class DutyAssignment {
        -id: uuid.UUID
        -title: string
        -officeSymbol: string
        -unit: Unit
        -assignee: *Member
    }

    class Member {
        -capid: uint
        -memberType: cap.MemberType
        -grade: cap.Grade
        -lastName: string
        -firstName: string
        -ageBracket: AB
    }

    class Unit {
        -id: uuid.UUID
        cap.Unit
    }
```


## User Module

# Entity Relationship Diagrams