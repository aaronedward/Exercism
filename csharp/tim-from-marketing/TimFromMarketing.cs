using System;

static class Badge
{
    public static string Print(int? id, string name, string? department)
    {
        if (id == null)
        {
            return $"{name} - {department.ToUpper()}";
        } else if (department == null && id == null)
        {
            return $"{name} - OWNER";
        } else {
            id ?? "OWNER";
            return $"[{id}] - {name} - {department.ToUpper()}";
        }
    }
}
