import { z } from "zod"

export const resourceIdSchema = z.enum(["water"])
export type ResourceId = z.infer<typeof resourceIdSchema>

export const resourceSchema = z.object({
  id: resourceIdSchema, 
  name: z.string(),
  amount: z.number(),
  delta: z.number(),
  total: z.number(),
  is_automated: z.boolean(),
})
export type Resource = z.infer<typeof resourceSchema>

export const toolIdSchema = z.enum(["search-water"])
export type ToolId = z.infer<typeof toolIdSchema>

const toolCostSchema = z.object({
  id: resourceIdSchema,
  name: z.string(),
  amount: z.number(),
})

export const toolSchema = z.object({
  id: toolIdSchema,
  name: z.string(),
  costs: z.array(toolCostSchema),
  is_enabled: z.boolean(),
})
export type Tool = z.infer<typeof toolSchema>

export const gameStateSchema = z.object({
  resources: z.array(resourceSchema),
  tools: z.array(toolSchema)
})
export type GameState = z.infer<typeof gameStateSchema>
